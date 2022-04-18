package service

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	repo "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/repository"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
)

type ProductService struct {
	repository *repo.Repository
}

func NewProductService(repo *repo.Repository) *ProductService {
	return &ProductService{repository: repo}
}

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

func (ps *ProductService) SearchProduct(keyword string, pag *helper.Pagination) (*dtos.ResponseAllProductsDto, error) {
	products, err := ps.repository.SearchProducts(keyword, pag)
	if err != nil {
		return nil, err
	}

	productsListResponse := helper.ConvertProductModelListToResponseProductDtoList(*products)
	return productsListResponse, nil
}

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

func (ps *ProductService) DeleteProduct(id string) error {

	// Response from repository
	err := ps.repository.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
