package models

import (
	"errors"
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"uniqueIndex"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"isAdmin" gorm:"default:false"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewString()
	return nil
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) IsValidPassword(password string) error {
	bytePassword := []byte(password)
	hashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(hashedPassword, bytePassword)
}

func (u *User) GenerateJwtToken(cfg *config.Config) string {
	jwtToken := jwt.New(jwt.SigningMethodHS512)

	jwtToken.Claims = jwt.MapClaims{
		"userId":   u.ID,
		"username": u.Username,
		"isAdmin":  u.IsAdmin,
		"exp":      time.Now().Add(time.Minute * time.Duration(cfg.JWTConfig.SessionTime)).Unix(),
	}

	token, _ := jwtToken.SignedString([]byte(cfg.JWTConfig.SecretKey))
	return token
}
