package main

import (
	"OrderPick/database"
	"OrderPick/middleware"
	"OrderPick/routes"

	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var ordercollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.ItemRoutes(router)
	routes.OrderRoutes(router)

	router.Run(":" + port)
}
