package controller

import (
	"OrderPick/helpers"
	"OrderPick/models"
	"OrderPick/repositories"
	"encoding/base64"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

type UserController struct {
	repo *repositories.UserRepository
}

func NewUserController(repo *repositories.UserRepository) *UserController {
	return &UserController{repo}
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	// Parse recordPerPage from query parameters
	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	// Get pagingState from query parameters (it should be base64 encoded)
	pagingStateBase64 := c.Query("pagingState")
	var pagingState []byte
	if pagingStateBase64 != "" {
		pagingState, err = base64.StdEncoding.DecodeString(pagingStateBase64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pagingState"})
			return
		}
	}

	// Call the repository to get users
	users, nextPageState, err := ctrl.repo.GetUsers(recordPerPage, pagingState)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing user items"})
		return
	}

	// Encode the nextPageState to base64 for the response
	var nextPageStateBase64 string
	if len(nextPageState) > 0 {
		nextPageStateBase64 = base64.StdEncoding.EncodeToString(nextPageState)
	}

	// Return the users and the nextPageState in the response
	c.JSON(http.StatusOK, gin.H{
		"users":         users,
		"nextPageState": nextPageStateBase64,
	})
}
func (ctrl *UserController) GetUser(c *gin.Context) {
	userId := c.Param("user_id")

	user, err := ctrl.repo.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while retrieving user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	// check user exit?
	foundUser, err := ctrl.repo.GetUserByEmail(*user.Email)
	if foundUser.User_id != "" && err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in fetching user by email",
			"detail": err})
		return
	}
	if foundUser.User_id != "" {
		c.JSON(http.StatusConflict, gin.H{"conflict": "email already registered"})
		return
	}

	hashedPassword := HashPassword(*user.Password)
	user.Password = &hashedPassword

	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.User_id = gocql.TimeUUID().String()

	token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id)
	user.Token = &token
	user.Refresh_Token = &refreshToken

	if err := ctrl.repo.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating user",
			"details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully",
		"token": user.Token})
}

func (ctrl *UserController) Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := ctrl.repo.GetUserByEmail(*user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found, login seems to be incorrect"})
		return
	}

	passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
	if !passwordIsValid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, foundUser.User_id)
	helpers.UpdateAllTokens(ctrl.repo, token, refreshToken, foundUser.User_id)

	c.JSON(http.StatusOK, foundUser)
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = "login or password is incorrect"
		check = false
	}
	return check, msg
}
