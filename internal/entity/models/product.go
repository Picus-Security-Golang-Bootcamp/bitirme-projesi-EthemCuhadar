package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product struct with relative fields
type Product struct {
	gorm.Model
	ID          string  `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"uniqueIndex; not null"`
	Description string  `json:"description" gorm:"not null"`
	Price       float64 `json:"price" gorm:"not null"`
	Stock       uint64  `json:"stock" gorm:"not null"`
	Brand       string  `json:"brand" gorm:"not null"`

	Category          []Category        `gorm:"many2many:products_categories;" json:"category"`
	ProductCategories []ProductCategory `gorm:"foreignkey:ProductId"`
}

// BeforeCreate sets a new uuid string to product ID
func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.NewString()
	return nil
}
