package service

import (
	"errors"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	repo "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/repository"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
)

// UserService struct with relative fields
type UserService struct {
	repository *repo.Repository
}

// NewCartService returns a new service structs
func NewUserService(repo *repo.Repository) *UserService {
	return &UserService{repository: repo}
}

// RegisterUser takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (us *UserService) RegisterUser(rr *dtos.RegisterRequest) (*dtos.RegisterResponse, error) {
	var registerResponse = dtos.RegisterResponse{}

	// Convert
	userModel := helper.ConvertRegisterRequestToUserModel(rr)

	// Response from repo
	_, err := us.repository.CreateUser(userModel)
	if err != nil {
		return nil, err
	}

	registerResponse.Message = &successRegister
	return &registerResponse, nil
}

// LoginUser takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (us *UserService) LoginUser(lr *dtos.LoginRequest, cfg *config.Config) (*dtos.LoginResponse, error) {

	// Response from repo
	userModel, err := us.repository.FindUser(*lr.Username)
	if err != nil {
		return nil, err
	}

	// Validation
	if userModel.IsValidPassword(*lr.Password) != nil {
		return nil, errors.New(failedLogin)
	}

	// Generate token
	token := userModel.GenerateJwtToken(cfg)
	return &dtos.LoginResponse{Jwt: &token}, nil
}

var successRegister = "Successful Registration"
var failedLogin = "Failed Login"
