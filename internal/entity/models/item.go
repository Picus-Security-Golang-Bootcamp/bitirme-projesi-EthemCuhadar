package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Item struct with relative fields
type Item struct {
	ID        string  `json:"id"`
	Price     float64 `json:"price"`
	Quantity  uint64  `json:"quantity"`
	CartId    string  `json:"cartId"`
	ProductId string  `json:"productId"`
}

// BeforeCreate sets a new uuid string to item ID
func (i *Item) BeforeCreate(tx *gorm.DB) error {
	i.ID = uuid.NewString()
	return nil
}
