package helper

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
)

// ConvertRequestCategoryDtoToCategoryModel simply takes the category request dto from
// client and convert into Category model which is used in the database.
func ConvertRequestCategoryDtoToCategoryModel(rcdto *dtos.RequestCategoryDto) models.Category {
	return models.Category{
		Name: *rcdto.Name,
	}
}

// ConvertCategoryModelToResponseCategoryDto simply takes a Category model from database
// and convert into Category dto model for client.
func ConvertCategoryModelToResponseCategoryDto(c *models.Category) *dtos.ResponseCategoryDto {
	return &dtos.ResponseCategoryDto{
		ID:   &c.ID,
		Name: &c.Name,
	}
}

// ConvertCategoryModelsToResponseAllCategoriesDto simply takes Category models from database
// and convert them into response Category dto model. Afterwards, it saves them into ResponseAllCategoryDto
// model for client.
func ConvertCategoryModelsToResponseAllCategoriesDto(c []models.Category) *dtos.ResponseAllCategoriesDto {
	categories := make([]*dtos.ResponseCategoryDto, 0)

	// Append categories
	for i := 0; i < len(c); i++ {
		var category = &dtos.ResponseCategoryDto{}
		category = ConvertCategoryModelToResponseCategoryDto(&c[i])
		categories = append(categories, category)
	}

	// Pass categories
	categoryListResponse := &dtos.ResponseAllCategoriesDto{
		Categories: categories,
	}
	return categoryListResponse
}
