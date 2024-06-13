package main

import (
	controller "OrderPick/controllers"
	"OrderPick/database"
	"OrderPick/kafka"
	"OrderPick/middleware"
	"OrderPick/repositories"
	"OrderPick/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setup database connection
	conn, err := database.SetupDBConnection()
	if err != nil {
		log.Fatal("Could not set up database connection:", err)
	}
	defer conn.Session.Close()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(conn.Session)
	itemRepo := repositories.NewItemRepository(conn.Session)
	orderRepo := repositories.NewOrderRepository(conn.Session)

	// Initialize controllers
	userController := controller.NewUserController(userRepo)
	itemController := controller.NewItemController(itemRepo)
	orderController := controller.NewOrderController(orderRepo)

	// Start Kafka consumer in a goroutine
	go kafka.ConsumeOrder(orderController)

	// Get port from environment variables or default to 8000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Initialize Gin router
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Adjust accordingly
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "token"}, // Include all headers you use
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Logger middleware
	router.Use(gin.Logger())

	// Define routes before authentication middleware
	routes.UserRoutes(router, userController)

	// Authentication middleware
	router.Use(middleware.Authentication())

	// Define routes after authentication middleware
	routes.ItemRoutes(router, itemController)
	routes.OrderRoutes(router, orderController)

	// Run the server
	router.Run(":" + port)
}
