package repo

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
)

type IProductRepository interface {
	CreateProduct(*models.Product) (*models.Product, error)
	FetchAllProducts() (*[]models.Product, error)
	FetchProduct(string) (*models.Product, error)
	FetchProductsOfSpecificCategory(string) (*[]models.Product, error)
}

func (r *Repository) CreateProduct(p *models.Product) (*models.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// DB query to create
	if err := r.DB.Create(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (r *Repository) FetchAllProducts(pag *helper.Pagination) (*[]models.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var products = &[]models.Product{}

	offset := (pag.Page) * pag.Limit
	queryBuilder := r.DB.Limit(int(pag.Limit)).Offset(int(offset)).Order(pag.Sort)

	// DB query to get
	if err := queryBuilder.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *Repository) FetchProduct(id string) (*models.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var product = &models.Product{}

	// DB query to get
	if err := r.DB.Preload("Category").Where(&models.Product{ID: id}).First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *Repository) FetchProductsOfSpecificCategory(id string, pag *helper.Pagination) (*[]models.Product, error) {

	// DB query to get
	category, err := r.FetchCategory(id)
	if err != nil {
		return nil, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	var products = &[]models.Product{}

	offset := (pag.Page) * pag.Limit
	queryBuilder := r.DB.Limit(int(pag.Limit)).Offset(int(offset)).Order(pag.Sort)

	// DB query
	queryBuilder.Model(&category).Association("Products").Find(&products)
	return products, nil
}

func (r *Repository) UpdateProduct(p *models.Product) (*models.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// DB query to update
	if err := r.DB.Save(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (r *Repository) DeleteProduct(id string) error {
	product, err := r.FetchProduct(id)
	if err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if result := r.DB.Unscoped().Delete(&product); result.Error != nil {
		return result.Error
	}
	return nil
}
