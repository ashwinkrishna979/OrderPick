package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID             primitive.ObjectID `bson:"_id"`
	Order_ID       string             `json:"order_id"`
	Item_id        *string            `json:"item_id" validate:"required"`
	Created_at     time.Time          `json:"created_at"`
	Packing_status bool               `json:"packing_status"`
}
