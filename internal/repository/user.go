package repo

import "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"

// CreateUser gets data from database and sends them into service, if there are no errors
func (r *Repository) CreateUser(u *models.User) (*models.User, error) {

	// DB query
	if err := r.DB.Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// FindUser gets data from database and sends them into service, if there are no errors
func (r *Repository) FindUser(username string) (*models.User, error) {
	var user = &models.User{}

	// DB query
	if err := r.DB.Where(&models.User{Username: username}).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
