package models

import (
	// "errors"
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User struct with relative fields
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

// BeforeCreate sets a new uuid string to user ID
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewString()
	return nil
}

// func (u *User) SetPassword(password string) error {
// 	if len(password) == 0 {
// 		return errors.New("password should not be empty")
// 	}
// 	bytePassword := []byte(password)
// 	hashedPassword, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
// 	u.Password = string(hashedPassword)
// 	return nil
// }

// IsValidPassword take a password and compares with hased one. It returns true
// if they mactch.
func (u *User) IsValidPassword(password string) error {
	bytePassword := []byte(password)
	hashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(hashedPassword, bytePassword)
}

// GenerateJwtToken takes configuration parameters for JWT token and creates a new
// JWT token string for logged in users. Afterwards, it returns the token.
func (u *User) GenerateJwtToken(cfg *config.Config) string {

	// New token
	jwtToken := jwt.New(jwt.SigningMethodHS512)

	// Claimes with relative user fields
	jwtToken.Claims = jwt.MapClaims{
		"userId":   u.ID,
		"username": u.Username,
		"isAdmin":  u.IsAdmin,
		"exp":      time.Now().Add(time.Minute * time.Duration(cfg.JWTConfig.SessionTime)).Unix(),
	}

	// Get the complete, signed token
	token, _ := jwtToken.SignedString([]byte(cfg.JWTConfig.SecretKey))
	return token
}
