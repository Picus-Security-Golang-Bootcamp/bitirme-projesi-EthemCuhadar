package service

import (
	"errors"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	repo "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/repository"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
)

type UserService struct {
	repository *repo.Repository
}

func NewUserService(repo *repo.Repository) *UserService {
	return &UserService{repository: repo}
}

func (us *UserService) RegisterUser(rr *dtos.RegisterRequest) (*dtos.RegisterResponse, error) {
	var registerResponse = dtos.RegisterResponse{}
	userModel := helper.ConvertRegisterRequestToUserModel(rr)
	_, err := us.repository.CreateUser(userModel)
	if err != nil {
		return nil, err
	}
	registerResponse.Message = &successRegister
	return &registerResponse, nil
}

func (us *UserService) LoginUser(lr *dtos.LoginRequest, cfg *config.Config) (*dtos.LoginResponse, error) {
	userModel, err := us.repository.FindUser(*lr.Username)
	if err != nil {
		return nil, err
	}
	if userModel.IsValidPassword(*lr.Password) != nil {
		return nil, errors.New(failedLogin)
	}
	token := userModel.GenerateJwtToken(cfg)
	return &dtos.LoginResponse{Jwt: &token}, nil
}

var successRegister = "Successful Registration"

// var successLogin = "Successful Login"
var failedLogin = "Failed Login"

// var failedRegister = "Failed Login"
