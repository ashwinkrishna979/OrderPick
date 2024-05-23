package controller

import (
	"OrderPick/database"
	"OrderPick/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var itemCollection *mongo.Collection = database.OpenCollection(database.Client, "item")
var validate = validator.New()

func GetItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := itemCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing the items"})
		}
		var allItems []bson.M
		if err = result.All(ctx, &allItems); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allItems)
	}
}

func GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		itemId := c.Param("item_id")
		var item models.Item
		err := itemCollection.FindOne(ctx, bson.M{"item_id": itemId}).Decode(&item)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the item"})
		}
		c.JSON(http.StatusOK, item)
	}
}

func CreateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var item models.Item
		defer cancel()

		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(item)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return

		}

		item.ID = primitive.NewObjectID()
		item.Item_id = item.ID.Hex()

		result, insertErr := itemCollection.InsertOne(ctx, item)
		if insertErr != nil {
			msg := fmt.Sprintf("Item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func UpdateItem() gin.HandlerFunc {
	return func(C *gin.Context) {

	}
}
