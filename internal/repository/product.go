package repo

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
	"go.uber.org/zap"
)

type IProductRepository interface {
	CreateProduct(*models.Product) (*models.Product, error)
	FetchAllProducts() (*[]models.Product, error)
	FetchProduct(string) (*models.Product, error)
	FetchProductsOfSpecificCategory(string) (*[]models.Product, error)
}

// CreateProduct gets data from database and sends them into service, if there are no errors
func (r *Repository) CreateProduct(p *models.Product) (*models.Product, error) {
	zap.L().Debug("repo.cart.CreateProduct")
	r.mu.Lock()
	defer r.mu.Unlock()

	// DB query to create
	if err := r.DB.Create(&p).Error; err != nil {
		zap.L().Debug("repo.cart.CreateProduct Error 1", zap.Reflect("error:", err))
		return nil, err
	}
	return p, nil
}

// FetchAllProducts gets data from database and sends them into service, if there are no errors
func (r *Repository) FetchAllProducts(pag *helper.Pagination) (*[]models.Product, error) {
	zap.L().Debug("repo.cart.FetchAllProducts")
	r.mu.Lock()
	defer r.mu.Unlock()
	var products = &[]models.Product{}

	// Pagination
	offset := (pag.Page) * pag.Limit
	queryBuilder := r.DB.Limit(int(pag.Limit)).Offset(int(offset)).Order(pag.Sort)

	// DB query to get
	if err := queryBuilder.Preload("Category").Find(&products).Error; err != nil {
		zap.L().Debug("repo.cart.FetchAllProducts Error 1", zap.Reflect("error:", err))
		return nil, err
	}
	return products, nil
}

// FetchProduct gets data from database and sends them into service, if there are no errors
func (r *Repository) FetchProduct(id string) (*models.Product, error) {
	zap.L().Debug("repo.cart.FetchProduct")
	r.mu.Lock()
	defer r.mu.Unlock()
	var product = &models.Product{}

	// DB query to get
	if err := r.DB.Preload("Category").Where(&models.Product{ID: id}).First(&product).Error; err != nil {
		zap.L().Debug("repo.cart.FetchProduct Error 1", zap.Reflect("error:", err))
		return nil, err
	}
	return product, nil
}

// SearchProducts gets data from database and sends them into service, if there are no errors
func (r *Repository) SearchProducts(keyword string, pag *helper.Pagination) (*[]models.Product, error) {
	zap.L().Debug("repo.cart.SearchProducts")
	r.mu.Lock()
	defer r.mu.Unlock()
	var products = &[]models.Product{}

	// Pagination
	offset := (pag.Page) * pag.Limit
	queryBuilder := r.DB.Limit(int(pag.Limit)).Offset(int(offset)).Order(pag.Sort)

	// DB query to get
	if err := queryBuilder.Where("name LIKE ?", "%"+keyword+"%").Find(&products).Error; err != nil {
		zap.L().Debug("repo.cart.SearchProducts Error 1", zap.Reflect("error:", err))
		return nil, err
	}
	return products, nil
}

// FetchProductsOfSpecificCategory gets data from database and sends them into service, if there are no errors
func (r *Repository) FetchProductsOfSpecificCategory(id string, pag *helper.Pagination) (*[]models.Product, error) {
	zap.L().Debug("repo.cart.FetchProductsOfSpecificCategory")
	// DB query to get
	category, err := r.FetchCategory(id)
	if err != nil {
		return nil, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	var products = &[]models.Product{}

	// Pagination
	offset := (pag.Page) * pag.Limit
	queryBuilder := r.DB.Limit(int(pag.Limit)).Offset(int(offset)).Order(pag.Sort)

	// DB query
	queryBuilder.Model(&category).Association("Products").Find(&products)
	return products, nil
}

// UpdateProduct gets data from database and sends them into service, if there are no errors
func (r *Repository) UpdateProduct(p *models.Product) (*models.Product, error) {
	zap.L().Debug("repo.cart.UpdateProduct")
	r.mu.Lock()
	defer r.mu.Unlock()

	// DB query to update
	if err := r.DB.Save(&p).Error; err != nil {
		zap.L().Debug("repo.cart.UpdateProduct Error 1", zap.Reflect("error:", err))
		return nil, err
	}
	return p, nil
}

// DeleteProduct gets data from database and sends them into service, if there are no errors
func (r *Repository) DeleteProduct(id string) error {
	zap.L().Debug("repo.cart.DeleteProduct")
	// DB query
	product, err := r.FetchProduct(id)
	if err != nil {
		zap.L().Debug("repo.cart.DeleteProduct Error 1", zap.Reflect("error:", err))
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	// DB query
	if result := r.DB.Unscoped().Delete(&product); result.Error != nil {
		zap.L().Debug("repo.cart.DeleteProduct Error 2", zap.Reflect("error:", err))
		return result.Error
	}
	return nil
}
