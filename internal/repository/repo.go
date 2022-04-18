package repo

import (
	"sync"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
	"gorm.io/gorm"
)

// Repository struct
type Repository struct {
	mu sync.Mutex
	DB *gorm.DB
}

// NewRepository returns a new repository structs
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

// Migrations migrate gorm models.
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
