package helper

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
	"golang.org/x/crypto/bcrypt"
)

func ConvertRegisterRequestToUserModel(rr *dtos.RegisterRequest) *models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte(*rr.Password), bcrypt.DefaultCost)
	return &models.User{
		FirstName: *rr.FirstName,
		LastName:  *rr.LastName,
		Username:  *rr.Username,
		Email:     *rr.Email,
		Password:  string(password),
		IsAdmin:   rr.IsAdmin,
	}
}
