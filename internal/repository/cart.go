package repo

import "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"

// CreateCart gets data from database and sends them into service, if there are no errors
func (r *Repository) CreateCart(c *models.Cart) (*models.Cart, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// DB query
	if err := r.DB.Create(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

// CreateCart gets data from database and sends them into service, if there are no errors
func (r *Repository) FetchCart(cart_id string) (*models.Cart, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var cart = &models.Cart{}

	// DB query
	if err := r.DB.Preload("Items").Where(&models.Cart{ID: cart_id}).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

// CreateCart gets data from database and sends them into service, if there are no errors
func (r *Repository) DeleteCart(cart_id string) error {
	cart, err := r.FetchCart(cart_id)
	if err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	// DB query
	for _, item := range cart.Items {
		r.DeleteItem(item.ID)
	}

	// DB query
	if result := r.DB.Delete(&cart); result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateCart gets data from database and sends them into service, if there are no errors
func (r *Repository) UpdateCart(c *models.Cart) (*models.Cart, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// DB query to update
	if err := r.DB.Save(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

// CreateCart gets data from database and sends them into service, if there are no errors
func (r *Repository) GetAllOrders(user_id string) (*[]models.Cart, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var carts = &[]models.Cart{}

	// DB query
	if err := r.DB.Preload("Items").Where("user_id = ? AND is_ordered = ?", user_id, true).Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}
