package service

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	repo "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/repository"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
)

// ProductService struct with relative fields
type ProductService struct {
	repository *repo.Repository
}

// NewCartService returns a new service structs
func NewProductService(repo *repo.Repository) *ProductService {
	return &ProductService{repository: repo}
}

// CreateProduct takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (ps *ProductService) CreateProduct(p *dtos.RequestProductDto) (*dtos.ResponseProductDto, error) {

	// Convert request into Model
	productModel := helper.ConvertRequestProductDtoToProductModel(p)

	// Response from repository
	product, err := ps.repository.CreateProduct(productModel)
	if err != nil {
		return nil, err
	}

	// Convert product into DTO
	productResponse := helper.ConvertProductModelToResponseProductDto(product)
	return productResponse, nil
}

// FetchAllProducts takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (ps *ProductService) FetchAllProducts(pag *helper.Pagination) (*dtos.ResponseAllProductsDto, error) {

	// Response from repository
	products, err := ps.repository.FetchAllProducts(pag)
	if err != nil {
		return nil, err
	}

	// Convert products into DTO
	productsListResponse := helper.ConvertProductModelListToResponseProductDtoList(*products)
	return productsListResponse, nil
}

// FetchProduct takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (ps *ProductService) FetchProduct(id string) (*dtos.ResponseProductDto, error) {

	// Response from repository
	product, err := ps.repository.FetchProduct(id)
	if err != nil {
		return nil, err
	}

	// Convert product into DTO
	productResponse := helper.ConvertProductModelToResponseProductDto(product)
	return productResponse, nil
}

// SearchProduct takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (ps *ProductService) SearchProduct(keyword string, pag *helper.Pagination) (*dtos.ResponseAllProductsDto, error) {

	// Response from repo
	products, err := ps.repository.SearchProducts(keyword, pag)
	if err != nil {
		return nil, err
	}

	productsListResponse := helper.ConvertProductModelListToResponseProductDtoList(*products)
	return productsListResponse, nil
}

// FetchProductsOfSpecificCategory takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (ps *ProductService) FetchProductsOfSpecificCategory(id string, pag *helper.Pagination) (*dtos.ResponseAllProductsDto, error) {

	// Response from repository
	products, err := ps.repository.FetchProductsOfSpecificCategory(id, pag)
	if err != nil {
		return nil, err
	}

	// Convert products into DTO
	productsListResponse := helper.ConvertProductModelListToResponseProductDtoList(*products)
	return productsListResponse, nil
}

// UpdateProduct takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (ps *ProductService) UpdateProduct(id string, p *dtos.RequestProductDto) (*dtos.ResponseProductDto, error) {

	// Convert request into Model
	productModel := helper.ConvertRequestProductDtoToProductModel(p)

	// Response from repository
	product, err := ps.repository.UpdateProduct(productModel)
	if err != nil {
		return nil, err
	}

	// Convert product into DTO
	productResponse := helper.ConvertProductModelToResponseProductDto(product)
	return productResponse, nil
}

// DeleteProduct takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (ps *ProductService) DeleteProduct(id string) error {

	// Response from repository
	err := ps.repository.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
