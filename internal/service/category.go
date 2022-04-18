package service

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
	repo "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/repository"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
)

// CategoryService is a struct which will be the bridge between handler and repository.
// It contains repository field.
type CategoryService struct {
	repository *repo.Repository
}

// NewCategoryService takes a repository and returns a new CategoryService.
func NewCategoryService(repo *repo.Repository) *CategoryService {
	return &CategoryService{repository: repo}
}

// CreateCategory Category model and pass the model to repository. Afterwards, it takes
// Category model from repository and convert into response Category dto model. Finally,
// returns the response Category dto model.
func (cs *CategoryService) CreateCategory(c *models.Category) (*dtos.ResponseCategoryDto, error) {

	// Response from repo
	category, err := cs.repository.CreateCategory(c)
	if err != nil {
		return nil, err
	}

	// Convert
	categoryResponse := helper.ConvertCategoryModelToResponseCategoryDto(category)
	return categoryResponse, nil
}

// FetchAllCategories takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (cs *CategoryService) FetchAllCategories(pag *helper.Pagination) (*dtos.ResponseAllCategoriesDto, error) {

	// Response from repo
	categories, err := cs.repository.FetchAllCategories(pag)
	if err != nil {
		return nil, err
	}

	// Convert
	categoryListResponse := helper.ConvertCategoryModelsToResponseAllCategoriesDto(*categories)
	if err != nil {
		return nil, err
	}
	return categoryListResponse, nil
}

// FetchCategory takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (cs *CategoryService) FetchCategory(id string) (*dtos.ResponseCategoryDto, error) {

	// Response from repo
	category, err := cs.repository.FetchCategory(id)
	if err != nil {
		return nil, err
	}

	// Convert
	categoryResponse := helper.ConvertCategoryModelToResponseCategoryDto(category)
	return categoryResponse, nil
}
