package handler

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/service"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/httpErrors"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// ProductHandler struct with relative fields
type ProductHandler struct {
	service *service.ProductService
	config  *config.Config
}

// NewProductHandler takes gin, service and config parameters and returns a new handler struct.
func NewProductHandler(r *gin.RouterGroup, service *service.ProductService, cfg *config.Config) {
	ph := &ProductHandler{
		service: service,
		config:  cfg,
	}

	// Endpoints with relative functions and methods

	// GET
	r.GET("/product/all", ph.FetchAllProducts)
	r.GET("/product/:product_id", ph.FetchProduct)
	r.GET("/category/:category_id/product", ph.FetchProductsOfSpecificCategory)
	r.GET("/product", ph.SearchProducts)

	// Middlewares
	r.Use(middleware.JWTMiddleware(cfg), middleware.AuthMiddleware(cfg))
	{
		r.POST("/product/create", ph.CreateProduct)
		r.PUT("/product/:product_id", ph.UpdateProduct)
		r.DELETE("/product/:product_id", ph.DeleteProduct)
	}

}

// CreateProduct handler takes informations from request body and bind them
// into struct. Afterwards, it validates and sends the struct to service. If the
// reponse is not nil, it serialize the data to response body
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

// FetchAllProducts serializes the data to response body, if response from
// service is not nil
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

// FetchProduct handler takes product id from url path and sends to service. If the
// reponse is not nil, it serialize the data to response body
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

// SearchProducts handler takes search keyword from url path and sends to service. If the
// reponse is not nil, it serialize the data to response body
func (ph *ProductHandler) SearchProducts(c *gin.Context) {
	keyword := c.Query("search")

	// Pagination
	pagination := helper.GeneratePaginationFromRequest(c)

	// Response from service
	products, err := ph.service.SearchProduct(keyword, pagination)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, products)
}

// FetchProductsOfSpecificCategory serializes the data to response body, if response from
// service is not nil
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

// UpdateProduct serializes the data to response body, if response from
// service is not nil
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

// DeleteProduct handler takes product id from url path and sends to service. If the
// reponse is not nil, it serialize the data to response body
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
