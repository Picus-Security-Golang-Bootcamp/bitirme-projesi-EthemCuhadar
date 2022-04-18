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

// CreateCategory gets data from database and sends them into service, if there are no errors
func (r *Repository) CreateCategory(c *models.Category) (*models.Category, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// DB query
	if err := r.DB.Create(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

// FetchAllCategories gets data from database and sends them into service, if there are no errors
func (r *Repository) FetchAllCategories(pag *helper.Pagination) (*[]models.Category, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var categories = &[]models.Category{}

	// Pagination
	offset := (pag.Page) * pag.Limit
	queryBuilder := r.DB.Limit(int(pag.Limit)).Offset(int(offset)).Order(pag.Sort)

	// DB query
	if err := queryBuilder.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// FetchCategory gets data from database and sends them into service, if there are no errors
func (r *Repository) FetchCategory(id string) (*models.Category, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var category = &models.Category{}

	// DB query
	if err := r.DB.Where(&models.Category{ID: id}).First(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
