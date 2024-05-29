package routes

import (
	controller "OrderPick/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine, orderController *controller.OrderController) {
	incomingRoutes.GET("/orders", orderController.GetOrders)
	incomingRoutes.GET("/orders/:order_id", orderController.GetOrder)
	incomingRoutes.POST("/orders", orderController.CreateOrder)
}
