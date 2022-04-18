package repo

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
	"go.uber.org/zap"
)

// FetchItem gets data from database and sends them into service, if there are no errors
func (r *Repository) FetchItem(item_id string) (*models.Item, error) {
	zap.L().Debug("repo.cart.FetchItem")
	r.mu.Lock()
	defer r.mu.Unlock()
	var item = &models.Item{}

	// DB query to get
	if err := r.DB.Where(&models.Item{ID: item_id}).First(&item).Error; err != nil {
		zap.L().Debug("repo.cart.FetchItem Error 1", zap.Reflect("error:", err))
		return nil, err
	}
	return item, nil
}

// UpdateItem gets data from database and sends them into service, if there are no errors
func (r *Repository) UpdateItem(i *models.Item) error {
	zap.L().Debug("repo.cart.UpdateItem")
	r.mu.Lock()
	defer r.mu.Unlock()

	// DB query to update
	if err := r.DB.Save(&i).Error; err != nil {
		zap.L().Debug("repo.cart.UpdateItem Error 1", zap.Reflect("error:", err))
		return err
	}
	return nil
}

// DeleteItem gets data from database and sends them into service, if there are no errors
func (r *Repository) DeleteItem(item_id string) error {
	zap.L().Debug("repo.cart.DeleteItem")

	// DB query
	item, err := r.FetchItem(item_id)
	if err != nil {
		zap.L().Debug("repo.cart.DeleteItem Error 1", zap.Reflect("error:", err))
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// DB query
	if result := r.DB.Unscoped().Delete(&item); result.Error != nil {
		zap.L().Debug("repo.cart.DeleteItem Error 2", zap.Reflect("error:", err))
		return result.Error
	}
	return nil
}

// FetchCategory gets data from database and sends them into service, if there are no errors
func (r *Repository) CreateItem(i *models.Item) (*models.Item, error) {
	zap.L().Debug("repo.cart.CreateItem")
	r.mu.Lock()
	defer r.mu.Unlock()

	// DB query
	if err := r.DB.Create(&i).Error; err != nil {
		zap.L().Debug("repo.cart.CreateItem Error 1", zap.Reflect("error:", err))
		return nil, err
	}
	return i, nil
}
