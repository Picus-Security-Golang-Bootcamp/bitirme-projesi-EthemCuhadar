package models

import (
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
