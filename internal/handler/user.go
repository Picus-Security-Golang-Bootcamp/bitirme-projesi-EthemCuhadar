package handler

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/service"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/httpErrors"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// UserHandler struct with relative fields
type UserHandler struct {
	service *service.UserService
	config  *config.Config
}

// NewProductHandler takes gin, service and config parameters and returns a new handler struct.
func NewUserHandler(r *gin.RouterGroup, service *service.UserService, cfg *config.Config) {
	uh := &UserHandler{
		service: service,
		config:  cfg,
	}

	// Endpoints with relative functions and methods

	// POST
	r.POST("/register", uh.RegisterUser)
	r.POST("/login", uh.LoginUser)
}

// RegisterUser handler takes informations from request body and bind them
// into struct. Afterwards, it validates and sends the struct to service. If the
// reponse is not nil, it serialize the data to response body
func (u *UserHandler) RegisterUser(c *gin.Context) {
	var registerRequest = &dtos.RegisterRequest{}

	// Binding
	if err := c.Bind(&registerRequest); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Response from service
	registerResponse, err := u.service.RegisterUser(registerRequest)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, registerResponse)
}

// RegisterUser handler takes informations from request body and bind them
// into struct. Afterwards, it validates and sends the struct to service. If the
// reponse is not nil, it serialize the data to response body
func (u *UserHandler) LoginUser(c *gin.Context) {
	var loginRequest = &dtos.LoginRequest{}

	// Bind
	if err := c.Bind(&loginRequest); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Validate
	if err := loginRequest.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Response from service
	loginResponse, err := u.service.LoginUser(loginRequest, u.config)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	// Serialization into response body
	c.JSON(http.StatusOK, loginResponse)
}
