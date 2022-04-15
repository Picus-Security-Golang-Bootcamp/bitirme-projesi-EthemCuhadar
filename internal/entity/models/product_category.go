package models

type ProductCategory struct {
	Category   *Category `gorm:"association_foreignkey:CategoryId"`
	CategoryId string
	Product    *Product `gorm:"association_foreignkey:ProductId"`
	ProductId  string
}

func (*ProductCategory) TableName() string {
	return "products_categories"
}
