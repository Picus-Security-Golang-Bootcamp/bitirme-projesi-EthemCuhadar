package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Category struct with relative fields
type Category struct {
	gorm.Model
	ID       string    `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" gorm:"uniqueIndex; not null"`
	Products []Product `gorm:"many2many:products_categories;"`
}

// BeforeCreate sets a new uuid string to category ID
func (c *Category) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.NewString()
	return nil
}
