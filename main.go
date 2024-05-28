package main

import (
	controller "OrderPick/controllers"
	"OrderPick/database"
	"OrderPick/middleware"
	"OrderPick/repositories"
	"OrderPick/routes"
	"log"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	conn, err := database.SetupDBConnection()
	if err != nil {
		log.Fatal("Could not set up database connection:", err)
	}
	defer conn.Session.Close()
	userRepo := repositories.NewUserRepository(conn.Session)
	itemRepo := repositories.NewItemRepository(conn.Session)
	userController := controller.NewUserController(userRepo)
	itemController := controller.NewItemController(itemRepo)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router, userController)
	router.Use(middleware.Authentication())

	routes.ItemRoutes(router, itemController)

	router.Run(":" + port)
}
