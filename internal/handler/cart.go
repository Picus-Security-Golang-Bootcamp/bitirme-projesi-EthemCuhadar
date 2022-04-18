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

// CartHandler struct with relative fields
type CartHandler struct {
	service *service.CartService
	config  *config.Config
}

// NewCartHandler takes gin, service and config parameters and returns a new handler struct.
func NewCartHandler(r *gin.RouterGroup, service *service.CartService, cfg *config.Config) {
	ch := &CartHandler{
		service: service,
		config:  cfg,
	}

	// Middlewares
	r.Use(middleware.JWTMiddleware(cfg), middleware.UserMiddleware(cfg))
	{
		// Endpoints with relative functions and methods

		// GET
		r.GET("/user/:user_id/order", ch.GetAllOrders)
		r.GET("user/:user_id/cart/:cart_id", ch.FetchCart)

		// POST
		r.POST("/user/:user_id/cart/create", ch.CreateCart)
		r.POST("user/:user_id/cart/:cart_id/complete", ch.CompleteOrder)
		r.POST("user/:user_id/cart/:cart_id", ch.AddItemToCart)

		// PUT
		r.PUT("user/:user_id/cart/:cart_id/item/:item_id", ch.UpdateItem)

		// DELETE
		r.DELETE("/user/:user_id/cart/:cart_id", ch.DeleteCart)
		r.DELETE("user/:user_id/cart/:cart_id/item/:item_id", ch.DeleteItem)
		r.DELETE("/user/:user_id/cart/:cart_id/cancel", ch.DeleteCart)

	}

}

// CreateCart handler takes cart informations from request body and bind them
// into struct. Afterwards, it validates and sends the struct to service. If the
// reponse is not nil, it serialize the data to response body
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

// DeleteCart handler takes cart id from url path and sends to service. If the
// reponse is not nil, it serialize the data to response body
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

// FetchCart handler takes cart id from url path and sends to service. If the
// reponse is not nil, it serialize the data to response body
func (ch *CartHandler) FetchCart(c *gin.Context) {
	cart_id := c.Param("cart_id")

	// Response from service
	cart, err := ch.service.FetchCart(cart_id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, cart)

}

// AddItemToCart handler takes item body item information from request body and bind them
// into struct. Afterwards, it validates and sends the struct to service. If the
// reponse is not nil, it serialize the data to response body
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

	// Response from service
	cartResponse, err := ch.service.AddItemToCart(&itemBody, cart_id)
	if err != nil {
		zap.L().Debug("handler.cart.additemtocart error 2")
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, cartResponse)
}

// UpdateItem handler takes item id and cart id from url path and sends to service. If the
// reponse is not nil, it serialize the data to response body
func (ch *CartHandler) UpdateItem(c *gin.Context) {
	cart_id := c.Param("cart_id")
	item_id := c.Param("item_id")
	itemBody := dtos.Item{}

	// Bind
	if err := c.Bind(&itemBody); err != nil {
		c.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
	}

	// Validation
	if err := itemBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	// Response from service
	cartResponse, err := ch.service.UpdateItem(&itemBody, item_id, cart_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	// Serialization into response body
	c.JSON(http.StatusOK, cartResponse)
}

// DeleteItem handler takes item id and cart id from url path and sends to service. If the
// reponse is not nil, it serialize the data to response body
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

// CompleteOrder takes cart id and sends to the service. If the
// reponse is not nil, it serialize the data to response body
func (ch *CartHandler) CompleteOrder(c *gin.Context) {
	cart_id := c.Param("cart_id")

	// Response from service
	cartResponse, err := ch.service.CompleteOrder(cart_id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, cartResponse)
}

// CancelOrder takes cart id and sends to the service. If the
// reponse is not nil, it serializes nil to response body.
func (ch *CartHandler) CancelOrder(c *gin.Context) {
	cart_id := c.Param("cart_id")

	// Response from service
	err := ch.service.CancelOrder(cart_id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, nil)
}

// GetAllOrders takes user id and sends to the service. If the
// reponse is not nil, it serialize the data to response body
func (ch *CartHandler) GetAllOrders(c *gin.Context) {
	user_id := c.Param("user_id")

	// Response from service
	orders, err := ch.service.GetAllOrders(user_id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, orders)
}
