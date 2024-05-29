package models

import (
	"time"

	"github.com/gocql/gocql"
)

type Order struct {
	Order_ID       gocql.UUID `json:"order_id"`
	Item_id        *string    `json:"item_id" validate:"required"`
	Created_at     time.Time  `json:"created_at"`
	Packing_status bool       `json:"packing_status"`
}
