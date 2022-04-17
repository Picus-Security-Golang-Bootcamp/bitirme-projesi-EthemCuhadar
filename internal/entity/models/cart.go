package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var minCartPrice float64 = 0

type Cart struct {
	gorm.Model
	ID         string    `json:"id"`
	UserID     string    `json:"userId"`
	Items      []Item    `json:"items"`
	Price      float64   `json:"price"`
	IsOrdered  bool      `json:"ordered" gorm:"default:false"`
	OrderTime  time.Time `json:"orderTime"`
	CancelTime time.Time `json:"cancelTime"`
}

type Item struct {
	ID       string  `json:"id"`
	Price    float64 `json:"price"`
	Quantity uint64  `json:"quantity"`
	CartId   string  `json:"cartId"`
	// OrderId  string  `json:"orderId"`
}

func (c *Cart) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.NewString()
	return nil
}

// func (c *Cart) CheckkOrder()(error){

// }

func (c *Cart) AddItem(i *Item) error {
	zap.L().Debug("models.cart.AddItem")

	_, item := c.SearchItemByID(i.ID)
	if item != nil {
		zap.L().Debug("Item already exists")
		return errors.New("Item already exists")
	}
	// item = &Item{
	// 	ID:       i.ID,
	// 	Price:    i.Price,
	// 	Quantity: i.Quantity,
	// }
	c.Items = append(c.Items, *i)
	return nil
}

func (c *Cart) UpdateItem(i *Item, quantity int) (err error) {
	if index, item := c.SearchItemByID(i.ID); index != -1 {
		item.Quantity = uint64(quantity)
	} else {
		return errors.New("Item not found")
	}
	return
}

func (c *Cart) RemoveItem(i *Item) error {
	if index, _ := c.SearchItemByID(i.ID); index != -1 {
		c.Items = append(c.Items[:index], c.Items[index+1:]...)
	}
	return errors.New("Item not found")
}

func (c *Cart) ValidateCart() error {
	totalPrice := CalculateTotalPrice(c)
	if totalPrice < minCartPrice {
		return errors.New("total cart price must be greater than min price")
	}
	c.Price = totalPrice
	return nil
}

func CalculateTotalPrice(c *Cart) (totalPrice float64) {
	for _, item := range c.Items {
		totalPrice += float64(item.Price) * float64(item.Quantity)
	}
	return
}

func (c *Cart) SearchItemByID(id string) (int, *Item) {
	for i, item := range c.Items {
		if item.ID == id {
			return i, &item
		}
	}
	return -1, nil
}
