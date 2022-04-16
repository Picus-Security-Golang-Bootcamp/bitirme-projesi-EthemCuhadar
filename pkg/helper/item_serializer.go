package helper

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
)

func ConvertItemModelToResponseItem(i *models.Item) *dtos.Item {
	return &dtos.Item{
		ID:       &i.ID,
		Price:    &i.Price,
		Quantity: &i.Quantity,
	}
}

func ConvertRequestItemToItemModel(i *dtos.Item) *models.Item {
	return &models.Item{
		ID:       *i.ID,
		Price:    *i.Price,
		Quantity: *i.Quantity,
	}
}
