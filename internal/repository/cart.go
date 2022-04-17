package repo

import "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"

func (r *Repository) CreateCart(c *models.Cart) (*models.Cart, error) {
	if err := r.DB.Create(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (r *Repository) FetchCart(cart_id string) (*models.Cart, error) {
	var cart = &models.Cart{}

	// DB query to get
	if err := r.DB.Preload("Items").Where(&models.Cart{ID: cart_id}).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (r *Repository) DeleteCart(cart_id string) error {
	cart, err := r.FetchCart(cart_id)
	if err != nil {
		return err
	}
	// for _, item := range cart.Items {
	// 	r.DeleteItem(item.ID)
	// }
	if result := r.DB.Delete(&cart); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) UpdateCart(c *models.Cart) (*models.Cart, error) {
	// DB query to update
	if err := r.DB.Save(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
