package repo

import "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"

func (r *Repository) CreateUser(u *models.User) (*models.User, error) {
	if err := r.DB.Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (r *Repository) FindUser(username string) (*models.User, error) {
	var user = &models.User{}
	if err := r.DB.Where(&models.User{Username: username}).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
