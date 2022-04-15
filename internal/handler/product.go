package handler

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/service"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/httpErrors"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

type ProductHandler struct {
	service *service.ProductService
	config  *config.Config
}

func NewProductHandler(r *gin.RouterGroup, service *service.ProductService, cfg *config.Config) {
	ph := &ProductHandler{
		service: service,
		config:  cfg,
	}

	r.GET("/product", ph.FetchAllProducts)
	r.GET("/product/:product_id", ph.FetchProduct)
	r.GET("/category/:category_id/product", ph.FetchProductsOfSpecificCategory)
	r.POST("/product/create", ph.CreateProduct)
	r.PUT("/product/:product_id", ph.UpdateProduct)
	r.DELETE("/product/:product_id", ph.DeleteProduct)

}

func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	productBody := dtos.RequestProductDto{}

	// Binding
	if err := c.Bind(&productBody); err != nil {
		c.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
	}

	// Validation
	if err := productBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	// Response from service
	product, err := ph.service.CreateProduct(&productBody)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) FetchAllProducts(c *gin.Context) {

	// Pagination
	pagination := helper.GeneratePaginationFromRequest(c)

	// Response from service
	products, err := ph.service.FetchAllProducts(pagination)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "message": err.Error()})
	}

	// Serialization into response body
	c.JSON(http.StatusOK, products)
}

func (ph *ProductHandler) FetchProduct(c *gin.Context) {
	product_id := c.Param("product_id")

	// Response from service
	product, err := ph.service.FetchProduct(product_id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) FetchProductsOfSpecificCategory(c *gin.Context) {
	category_id := c.Param("category_id")

	// Pagination
	pagination := helper.GeneratePaginationFromRequest(c)

	// Response from service
	products, err := ph.service.FetchProductsOfSpecificCategory(category_id, pagination)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, products)
}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	category_id := c.Param("category_id")
	productBody := dtos.RequestProductDto{ID: &category_id}

	// Binding
	if err := c.Bind(&productBody); err != nil {
		c.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
	}

	// Validation
	if err := productBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	// Response from service
	product, err := ph.service.UpdateProduct(category_id, &productBody)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {
	product_id := c.Param("product_id")

	// Response from service
	err := ph.service.DeleteProduct(product_id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusNoContent, nil)
}
