package handler

import (
	"fmt"
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/service"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/httpErrors"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// CategoryHandler struct with relative fields
type CategoryHandler struct {
	service *service.CategoryService
	config  *config.Config
}

// NewCategoryHandler takes gin, service and config parameters and returns a new handler struct.
func NewCategoryHandler(r *gin.RouterGroup, service *service.CategoryService, cfg *config.Config) {
	ch := &CategoryHandler{
		service: service,
		config:  cfg,
	}

	// Endpoints with relative functions and methods

	//GET
	r.GET("/", ch.FetchAllCategories)
	r.GET("/:category_id", ch.FetchCategory)

	// Middleware
	r.Use(middleware.AuthMiddleware(cfg), middleware.JWTMiddleware(cfg))
	{
		// POST
		r.POST("/create", ch.CreateCategory)
	}
}

// CreateCategory takes csv file from request body and read it. Afterwards,
// it validates and sends the struct to service. If the
// reponse is not nil, it serialize the data to response body
func (ch *CategoryHandler) CreateCategory(c *gin.Context) {

	// CSV file
	csvPartFile, _, openErr := c.Request.FormFile("file")
	if openErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("file err : %s", openErr.Error()))
		return
	}

	// Read CSV
	categoryBody, err := helper.ReadCSV(csvPartFile)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Response from service
	category, err := ch.service.CreateCategory(categoryBody)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, category)
}

// FetchAllCategories serializes the data to response body, if response from
// service is not nil
func (ch *CategoryHandler) FetchAllCategories(c *gin.Context) {

	// Generate Pagination
	pagination := helper.GeneratePaginationFromRequest(c)

	// Response from service
	categories, err := ch.service.FetchAllCategories(pagination)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, categories)
}

// FetchCategory takes category id from url path and sends to service. If the
// reponse is not nil, it serialize the data to response body
func (ch *CategoryHandler) FetchCategory(c *gin.Context) {
	id := c.Param("category_id")

	// Response from service
	category, err := ch.service.FetchCategory(id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, category)
}
