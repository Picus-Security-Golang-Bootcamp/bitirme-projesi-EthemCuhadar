package helper

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
	"github.com/go-openapi/strfmt"
)

func ConvertCreateCartRequestToCartModel(c *dtos.CreateCartRequest) *models.Cart {
	return &models.Cart{
		UserID: *c.UserID,
	}
}

func ConvertCartModelToCreateCartResponse(c *models.Cart) *dtos.CreateCartResponse {
	items := make([]*dtos.Item, 0)

	// Append categories
	for i := 0; i < len(c.Items); i++ {
		var item = &dtos.Item{}
		item = ConvertItemModelToResponseItem(&c.Items[i])
		items = append(items, item)
	}

	return &dtos.CreateCartResponse{
		UserID:     &c.UserID,
		ID:         &c.ID,
		Price:      &c.Price,
		IsOrdered:  &c.IsOrdered,
		OrderTime:  (*strfmt.Date)(&c.OrderTime),
		CancelTime: (*strfmt.Date)(&c.CancelTime),
		Item:       items,
	}
}
