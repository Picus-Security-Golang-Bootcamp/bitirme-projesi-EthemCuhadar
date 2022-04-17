package repo

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Migrations() {
	productPrototype := models.Product{}
	categoryPrototype := models.Category{}
	userPrototype := models.User{}
	cartPrototype := models.Cart{}
	itemPrototype := models.Item{}
	r.DB.AutoMigrate(
		&productPrototype,
		&categoryPrototype,
		&userPrototype,
		&cartPrototype,
		&itemPrototype,
	)
}
