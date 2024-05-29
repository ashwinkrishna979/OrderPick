package controller

import (
	"OrderPick/models"
	"OrderPick/repositories"
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type OrderController struct {
	repo *repositories.OrderRepository
}

func NewOrderController(repo *repositories.OrderRepository) *OrderController {
	return &OrderController{repo}
}

func (ctrl *OrderController) GetOrders(c *gin.Context) {
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

	// Call the repository to get orders
	orders, nextPageState, err := ctrl.repo.GetOrders(recordPerPage, pagingState)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing orders"})
		return
	}

	// Encode the nextPageState to base64 for the response
	var nextPageStateBase64 string
	if len(nextPageState) > 0 {
		nextPageStateBase64 = base64.StdEncoding.EncodeToString(nextPageState)
	}

	// Return the orders and the nextPageState in the response
	c.JSON(http.StatusOK, gin.H{
		"orders":        orders,
		"nextPageState": nextPageStateBase64,
	})
}

func (ctrl *OrderController) GetOrder(c *gin.Context) {
	orderId := c.Param("order_id")

	order, err := ctrl.repo.GetOrderById(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while retrieving order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(order)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	order.Order_ID = gocql.TimeUUID()
	order.Created_at = time.Now()
	order.Packing_status = false

	if err := ctrl.repo.CreateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating order",
			"details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order created successfully"})
}
