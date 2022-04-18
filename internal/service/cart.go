package service

import (
	"errors"
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	repo "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/repository"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/httpErrors"
	"go.uber.org/zap"
)

// CartHandler struct with relative fields
type CartService struct {
	repository *repo.Repository
}

// NewCartService returns a new service structs
func NewCartService(repo *repo.Repository) *CartService {
	return &CartService{repository: repo}
}

// CreateCart takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (cs *CartService) CreateCart(c *dtos.CreateCartRequest) (*dtos.CreateCartResponse, error) {

	// Convert
	cartRequest := helper.ConvertCreateCartRequestToCartModel(c)

	// Response from Repo
	cartModel, err := cs.repository.CreateCart(cartRequest)
	if err != nil {
		return nil, err
	}

	// Convert product into DTO
	cartResponse := helper.ConvertCartModelToCreateCartResponse(cartModel)
	return cartResponse, nil
}

// DeleteCart takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (cs *CartService) DeleteCart(cart_id string) error {

	// Response from Repo
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return err
	}

	// Check Ordered
	if cartModel.IsOrdered {
		if helper.InTimeSpan(cartModel.OrderTime, cartModel.CancelTime, time.Now()) {
			err := cs.repository.DeleteCart(cart_id)
			if err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("cancel time passed")
		}
	} else {

		// Response from Repo
		err := cs.repository.DeleteCart(cart_id)
		if err != nil {
			return err
		}
	}
	return nil
}

// FetchCart takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (cs *CartService) FetchCart(cart_id string) (*dtos.CreateCartResponse, error) {

	// Response from Repo
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return nil, err
	}

	if err := cartModel.ValidateCart(); err != nil {
		return nil, err
	}

	// Convert
	cartResponse := helper.ConvertCartModelToCreateCartResponse(cartModel)
	return cartResponse, nil
}

// AddItemToCart takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (cs *CartService) AddItemToCart(item *dtos.Item, cart_id string) (*dtos.CreateCartResponse, error) {

	// Response from Repo
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		zap.L().Debug("handler.service.additemtocart error 2")
		return nil, err
	}

	if cartModel.IsOrdered {
		return nil, errors.New("items ordered")
	}

	// Response from Repo
	product, err := cs.repository.FetchProduct(*item.ProductID)
	if err != nil {
		return nil, err
	}

	item.Price = product.Price

	product.Stock -= *item.Quantity

	// Response from Repo
	if _, err := cs.repository.UpdateProduct(product); err != nil {
		return nil, err
	}

	// Convert
	itemModel := helper.ConvertRequestItemToItemModel(item)
	itemModel.CartId = cart_id

	// Response from Repo
	if _, err := cs.repository.CreateItem(itemModel); err != nil {
		zap.L().Debug("handler.service.additemtocart error 1")
		return nil, err
	}

	// Response from Repo
	updatedCartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		zap.L().Debug("handler.service.additemtocart error 2")
		return nil, err
	}

	if err := updatedCartModel.ValidateCart(); err != nil {
		zap.L().Debug("handler.service.additemtocart error 3")
		return nil, err
	}

	// Convert
	cartResponse := helper.ConvertCartModelToCreateCartResponse(updatedCartModel)
	return cartResponse, nil
}

// UpdateItem takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (cs *CartService) UpdateItem(item *dtos.Item, item_id string, cart_id string) (*dtos.CreateCartResponse, error) {

	// Response from Repo
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return nil, err
	}
	if cartModel.IsOrdered {
		return nil, errors.New("items ordered")
	}

	// Response from Repo
	oldItem, err := cs.repository.FetchItem(item_id)
	if err != nil {
		return nil, err
	}

	// Response from Repo
	product, err := cs.repository.FetchProduct(*item.ProductID)
	if err != nil {
		return nil, err
	}

	product.Stock -= *item.Quantity - oldItem.Quantity

	// Response from Repo
	if _, err := cs.repository.UpdateProduct(product); err != nil {
		return nil, err
	}

	oldItem.Quantity = *item.Quantity

	// Response from Repo
	if err := cs.repository.UpdateItem(oldItem); err != nil {
		return nil, err
	}

	// Response from Repo
	updatedCart, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return nil, err
	}

	// Validate
	if err := updatedCart.ValidateCart(); err != nil {
		return nil, err
	}

	// Convert
	cartResponse := helper.ConvertCartModelToCreateCartResponse(updatedCart)
	return cartResponse, nil
}

// DeleteItem takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (cs *CartService) DeleteItem(item_id string, cart_id string) (*dtos.CreateCartResponse, error) {

	// Response from Repo
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return nil, err
	}
	if cartModel.IsOrdered {
		return nil, httpErrors.OrderError
	}

	// Response from Repo
	item, err := cs.repository.FetchItem(item_id)
	if err != nil {
		return nil, err
	}

	// Response from Repo
	product, err := cs.repository.FetchProduct(item.ProductId)
	if err != nil {
		return nil, err
	}

	product.Stock += item.Quantity

	// Response from Repo
	if _, err := cs.repository.UpdateProduct(product); err != nil {
		return nil, err
	}

	// Response from Repo
	if err := cs.repository.DeleteItem(item_id); err != nil {
		return nil, err
	}

	// Validate
	if err := cartModel.ValidateCart(); err != nil {
		return nil, err
	}

	// Convert
	cartResponse := helper.ConvertCartModelToCreateCartResponse(cartModel)
	return cartResponse, nil
}

// CompleteOrder takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (cs *CartService) CompleteOrder(cart_id string) (*dtos.CreateCartResponse, error) {

	// Response from Repo
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return nil, err
	}
	if !cartModel.IsOrdered {
		cartModel.IsOrdered = true
		cartModel.OrderTime = time.Now()
		cartModel.CancelTime = time.Now().Add(336 * time.Hour)
	} else {
		return nil, errors.New("aldready ordered")
	}

	// Response from Repo
	updatedCartModel, err := cs.repository.UpdateCart(cartModel)
	if err != nil {
		return nil, err
	}

	// Validate
	if err := updatedCartModel.ValidateCart(); err != nil {
		return nil, err
	}

	// Convert
	cartResponse := helper.ConvertCartModelToCreateCartResponse(updatedCartModel)
	return cartResponse, nil
}

// CancelOrder takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (cs *CartService) CancelOrder(cart_id string) error {

	// Response from Repo
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return err
	}
	if cartModel.IsOrdered {
		if helper.InTimeSpan(cartModel.OrderTime, cartModel.CancelTime, time.Now()) {
			err := cs.repository.DeleteCart(cart_id)
			if err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("cancel time passed")
		}
	} else {
		err := cs.repository.DeleteCart(cart_id)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAllOrders takes data from handler and returns response data if there is no error.
// Otherwise, it returns nil and error
func (cs *CartService) GetAllOrders(user_id string) ([]dtos.CreateCartResponse, error) {

	// Response from Repo
	cartModels, err := cs.repository.GetAllOrders(user_id)
	cartResponses := []dtos.CreateCartResponse{}
	if err != nil {
		return nil, err
	}

	for _, cart := range *cartModels {
		err := cart.ValidateCart()
		if err != nil {
			return nil, err
		}
		cartResponse := helper.ConvertCartModelToCreateCartResponse(&cart)
		cartResponses = append(cartResponses, *cartResponse)
	}
	return cartResponses, err
}
