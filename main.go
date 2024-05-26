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
	repo := repositories.NewUserRepository(conn.Session)
	userController := controller.NewUserController(repo)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router, userController)
	router.Use(middleware.Authentication())

	router.Run(":" + port)
}
