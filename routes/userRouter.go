package routes

import (
	controller "OrderPick/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine, userController *controller.UserController) {
	incomingRoutes.GET("/users", userController.GetUsers)
	incomingRoutes.GET("/users/:user_id", userController.GetUser)
	incomingRoutes.POST("/users/signup", userController.SignUp)
	incomingRoutes.POST("/users/login", userController.Login)
}
