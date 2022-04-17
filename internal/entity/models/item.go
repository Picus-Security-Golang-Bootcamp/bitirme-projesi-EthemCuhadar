package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID        string  `json:"id"`
	Price     float64 `json:"price"`
	Quantity  uint64  `json:"quantity"`
	CartId    string  `json:"cartId"`
	ProductId string  `json:"productId"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) error {
	i.ID = uuid.NewString()
	return nil
}
