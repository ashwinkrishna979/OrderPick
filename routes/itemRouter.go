package routes

import (
	controller "OrderPick/controllers"

	"github.com/gin-gonic/gin"
)

func ItemRoutes(incomingRoutes *gin.Engine, itemController *controller.ItemController) {

	incomingRoutes.GET("/items", itemController.GetItems)
	incomingRoutes.GET("/items/:item_id", itemController.GetItem)
	incomingRoutes.POST("/items", itemController.CreateItem)
}
