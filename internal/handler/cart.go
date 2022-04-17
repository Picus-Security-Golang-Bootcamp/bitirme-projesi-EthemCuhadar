package handler

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/service"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/httpErrors"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"go.uber.org/zap"
)

type CartHandler struct {
	service *service.CartService
	config  *config.Config
}

func NewCartHandler(r *gin.RouterGroup, service *service.CartService, cfg *config.Config) {
	ch := &CartHandler{
		service: service,
		config:  cfg,
	}
	r.Use(middleware.JWTMiddleware(cfg), middleware.UserMiddleware(cfg))
	{
		r.POST("/user/:user_id/cart/create", ch.CreateCart)
		r.DELETE("/user/:user_id/cart/:cart_id", ch.DeleteCart)
		r.GET("user/:user_id/cart/:cart_id", ch.FetchCart)
		r.POST("user/:user_id/cart/:cart_id", ch.AddItemToCart)
		r.POST("user/:user_id/cart/:cart_id/item/:item_id", ch.UpdateItem)
		r.DELETE("user/:user_id/cart/:cart_id/item/:item_id", ch.DeleteItem)
		r.POST("user/:user_id/cart/:cart_id/complete", ch.CompleteOrder)
		r.DELETE("/user/:user_id/cart/:cart_id/cancel", ch.DeleteCart)
	}

}

func (ch *CartHandler) CreateCart(c *gin.Context) {
	cartBody := dtos.CreateCartRequest{}

	// Binding
	if err := c.Bind(&cartBody); err != nil {
		c.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
	}

	// Validation
	if err := cartBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	// Response from service
	cartResponse, err := ch.service.CreateCart(&cartBody)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, cartResponse)
}

func (ch *CartHandler) DeleteCart(c *gin.Context) {
	cart_id := c.Param("cart_id")

	// Response from service
	err := ch.service.DeleteCart(cart_id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusNoContent, nil)
}

func (ch *CartHandler) FetchCart(c *gin.Context) {
	cart_id := c.Param("cart_id")

	cart, err := ch.service.FetchCart(cart_id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, cart)

}

func (ch *CartHandler) AddItemToCart(c *gin.Context) {
	itemBody := dtos.Item{}
	cart_id := c.Param("cart_id")

	zap.L().Debug("handler.cart.additemtocart")

	// Binding
	if err := c.Bind(&itemBody); err != nil {
		zap.L().Debug("handler.cart.additemtocart error 1")
		c.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
	}

	// Validation
	if err := itemBody.Validate(strfmt.NewFormats()); err != nil {
		zap.L().Debug("handler.cart.additemtocart error 2")
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	cartResponse, err := ch.service.AddItemToCart(&itemBody, cart_id)
	if err != nil {
		zap.L().Debug("handler.cart.additemtocart error 2")
		c.JSON(httpErrors.ErrorResponse(err))
	}

	c.JSON(http.StatusOK, cartResponse)
}

func (ch *CartHandler) UpdateItem(c *gin.Context) {
	cart_id := c.Param("cart_id")

	itemBody := dtos.Item{}

	if err := c.Bind(&itemBody); err != nil {
		c.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
	}

	// Validation
	if err := itemBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	cartResponse, err := ch.service.UpdateItem(&itemBody, cart_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, cartResponse)
}

func (ch *CartHandler) DeleteItem(c *gin.Context) {
	item_id := c.Param("item_id")
	cart_id := c.Param("cart_id")

	// Response from service
	cartResponse, err := ch.service.DeleteItem(item_id, cart_id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, cartResponse)
}

func (ch *CartHandler) CompleteOrder(c *gin.Context) {

	cart_id := c.Param("cart_id")
	cartResponse, err := ch.service.CompleteOrder(cart_id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}
	c.JSON(http.StatusOK, cartResponse)
}

func (ch *CartHandler) CancelOrder(c *gin.Context) {
	cart_id := c.Param("cart_id")
	err := ch.service.CancelOrder(cart_id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}
	c.JSON(http.StatusOK, nil)
}
