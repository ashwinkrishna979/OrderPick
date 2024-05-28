package controller

import (
	"OrderPick/models"
	"OrderPick/repositories"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type ItemController struct {
	repo *repositories.ItemRepository
}

func NewItemController(repo *repositories.ItemRepository) *ItemController {
	return &ItemController{repo}
}

func (ctrl *ItemController) GetItems(c *gin.Context) {
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
	items, nextPageState, err := ctrl.repo.GetItems(recordPerPage, pagingState)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing items"})
		return
	}

	// Encode the nextPageState to base64 for the response
	var nextPageStateBase64 string
	if len(nextPageState) > 0 {
		nextPageStateBase64 = base64.StdEncoding.EncodeToString(nextPageState)
	}

	// Return the users and the nextPageState in the response
	c.JSON(http.StatusOK, gin.H{
		"items":         items,
		"nextPageState": nextPageStateBase64,
	})
}

func (ctrl *ItemController) GetItem(c *gin.Context) {
	itemId := c.Param("item_id")

	item, err := ctrl.repo.GetItemById(itemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while retrieving item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (ctrl *ItemController) CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validattionErr := validate.Struct(item)
	if validattionErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validattionErr.Error()})
		return
	}
	item.Item_id = gocql.TimeUUID().String()

	if err := ctrl.repo.CreateItem(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating item",
			"details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "item created successfully"})
}
