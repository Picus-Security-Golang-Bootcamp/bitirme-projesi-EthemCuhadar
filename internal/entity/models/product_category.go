package models

// ProductCategory gorm model for many2many relationship between
// product and category models.
type ProductCategory struct {
	Category   *Category `gorm:"association_foreignkey:CategoryId"`
	CategoryId string
	Product    *Product `gorm:"association_foreignkey:ProductId"`
	ProductId  string
}

// TableName sets a new table of for ProductCategory model in database.
func (*ProductCategory) TableName() string {
	return "products_categories"
}
