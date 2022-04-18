package helper

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
)

// ConvertItemModelToResponseItem
func ConvertItemModelToResponseItem(i *models.Item) *dtos.Item {
	return &dtos.Item{
		ProductID: &i.ProductId,
		Price:     i.Price,
		Quantity:  &i.Quantity,
	}
}

// ConvertRequestItemToItemModel
func ConvertRequestItemToItemModel(i *dtos.Item) *models.Item {
	return &models.Item{
		ProductId: *i.ProductID,
		Price:     i.Price,
		Quantity:  *i.Quantity,
	}
}
