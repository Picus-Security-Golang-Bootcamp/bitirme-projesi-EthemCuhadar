package handler

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/service"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/httpErrors"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(r *gin.RouterGroup, service *service.CategoryService) {
	ch := &CategoryHandler{service: service}

	r.POST("/create", ch.CreateCategory)
	r.GET("/", ch.FetchAllCategories)
	r.GET("/:category_id", ch.FetchCategory)
}

func (ch *CategoryHandler) CreateCategory(c *gin.Context) {
	categoryBody := models.Category{}
	if err := c.Bind(&categoryBody); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}
	category, err := ch.service.CreateCategory(&categoryBody)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}
	c.JSON(http.StatusOK, category)
}

func (ch *CategoryHandler) FetchAllCategories(c *gin.Context) {

	pagination := helper.GeneratePaginationFromRequest(c)

	categories, err := ch.service.FetchAllCategories(pagination)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	c.JSON(http.StatusOK, categories)
}

func (ch *CategoryHandler) FetchCategory(c *gin.Context) {
	id := c.Param("category_id")
	category, err := ch.service.FetchCategory(id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}
	c.JSON(http.StatusOK, category)
}
