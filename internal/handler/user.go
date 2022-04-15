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

type UserHandler struct {
	service *service.UserService
	config  *config.Config
}

func NewUserHandler(r *gin.RouterGroup, service *service.UserService, cfg *config.Config) {
	uh := &UserHandler{
		service: service,
		config:  cfg,
	}

	r.POST("/register", uh.RegisterUser)
	r.POST("/login", uh.LoginUser)
}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	var registerRequest = &dtos.RegisterRequest{}
	if err := c.Bind(&registerRequest); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}
	registerResponse, err := u.service.RegisterUser(registerRequest)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}
	c.JSON(http.StatusOK, registerResponse)
}

func (u *UserHandler) LoginUser(c *gin.Context) {
	var loginRequest = &dtos.LoginRequest{}
	if err := c.Bind(&loginRequest); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	if err := loginRequest.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	loginResponse, err := u.service.LoginUser(loginRequest, u.config)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}
	c.JSON(http.StatusOK, loginResponse)
}
