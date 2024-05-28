package models

type Item struct {
	Name    *string `json:"name" validate:"required,min=2,max=100"`
	Item_id string  `json:"item_id"`
}
