package repo

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
)

type ICategoryRepository interface {
	CreateCategory(*models.Category) (*models.Category, error)
	FetchAllCategories() (*[]models.Category, error)
	FetchCategory(string) (*models.Category, error)
}

func (r *Repository) CreateCategory(c *models.Category) (*models.Category, error) {
	if err := r.DB.Create(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (r *Repository) FetchAllCategories(pag *helper.Pagination) (*[]models.Category, error) {
	var categories = &[]models.Category{}

	offset := (pag.Page) * pag.Limit
	queryBuilder := r.DB.Limit(int(pag.Limit)).Offset(int(offset)).Order(pag.Sort)

	if err := queryBuilder.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *Repository) FetchCategory(id string) (*models.Category, error) {
	var category = &models.Category{}
	if err := r.DB.Where(&models.Category{ID: id}).First(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
