package helper

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
)

// ConvertRequestProductDtoToProductModel take request Product dto and convert into
// Product model which are going to be used in database. Categories are going to be converted
// and saved in "Categories" field in the product model.
func ConvertRequestProductDtoToProductModel(rpdto *dtos.RequestProductDto) *models.Product {
	categories := make([]models.Category, 0)

	// Append categories
	for _, c := range rpdto.Categories {
		categories = append(categories, ConvertRequestCategoryDtoToCategoryModel(c))
	}

	return &models.Product{
		ID:          *rpdto.ID,
		Name:        *rpdto.Name,
		Category:    categories,
		Description: *rpdto.Description,
		Price:       *rpdto.Price,
		Stock:       *rpdto.Stock,
		Brand:       *rpdto.Brand,
	}
}

// ConvertProductModelToResponseProductDto simply takes Product model and convert into
// response Product dto model for client.
func ConvertProductModelToResponseProductDto(p *models.Product) *dtos.ResponseProductDto {
	return &dtos.ResponseProductDto{
		ID:          &p.ID,
		Name:        &p.Name,
		Description: &p.Description,
		Price:       &p.Price,
		Stock:       &p.Stock,
		Brand:       &p.Brand,
	}
}

// ConvertProductModelListToResponseProductDtoList takes Product models and convert them
// into response Product dto model. Afterwards, it saves them in "Products" field.
func ConvertProductModelListToResponseProductDtoList(ps []models.Product) *dtos.ResponseAllProductsDto {
	products := make([]*dtos.ResponseProductDto, 0)

	// Append
	for i := 0; i < len(ps); i++ {
		var product = &dtos.ResponseProductDto{}
		product = ConvertProductModelToResponseProductDto(&ps[i])
		products = append(products, product)
	}

	productListResponse := &dtos.ResponseAllProductsDto{
		Products: products,
	}
	return productListResponse
}
