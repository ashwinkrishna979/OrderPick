package routes

import (
	controller "OrderPick/controllers"

	"github.com/gin-gonic/gin"
)

func ItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/items", controller.GetItems())
	incomingRoutes.GET("/items/:item_id", controller.GetItem())
	incomingRoutes.POST("/items", controller.CreateItem())
	incomingRoutes.PATCH("/items/:item_id", controller.UpdateItem())
}
