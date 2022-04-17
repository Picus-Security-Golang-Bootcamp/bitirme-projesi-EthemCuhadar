package service

import (
	"errors"
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	repo "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/repository"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
	"go.uber.org/zap"
)

type CartService struct {
	repository *repo.Repository
}

func NewCartService(repo *repo.Repository) *CartService {
	return &CartService{repository: repo}
}

func (cs *CartService) CreateCart(c *dtos.CreateCartRequest) (*dtos.CreateCartResponse, error) {
	cartRequest := helper.ConvertCreateCartRequestToCartModel(c)

	cartModel, err := cs.repository.CreateCart(cartRequest)
	if err != nil {
		return nil, err
	}

	// Convert product into DTO
	cartResponse := helper.ConvertCartModelToCreateCartResponse(cartModel)
	return cartResponse, nil
}

func (cs *CartService) DeleteCart(cart_id string) error {
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

func (cs *CartService) FetchCart(cart_id string) (*dtos.CreateCartResponse, error) {
	// Response from repository
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return nil, err
	}

	if err := cartModel.ValidateCart(); err != nil {
		return nil, err
	}

	// Convert product into DTO
	cartResponse := helper.ConvertCartModelToCreateCartResponse(cartModel)
	return cartResponse, nil
}

func (cs *CartService) AddItemToCart(item *dtos.Item, cart_id string) (*dtos.CreateCartResponse, error) {

	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		zap.L().Debug("handler.service.additemtocart error 2")
		return nil, err
	}

	if cartModel.IsOrdered {
		return nil, errors.New("items ordered")
	}

	zap.L().Debug("handler.cart.additemtocart")
	itemModel := helper.ConvertRequestItemToItemModel(item)
	itemModel.CartId = cart_id

	if _, err := cs.repository.CreateItem(itemModel); err != nil {
		zap.L().Debug("handler.service.additemtocart error 1")
		return nil, err
	}

	updatedCartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		zap.L().Debug("handler.service.additemtocart error 2")
		return nil, err
	}

	if err := updatedCartModel.ValidateCart(); err != nil {
		zap.L().Debug("handler.service.additemtocart error 3")
		return nil, err
	}

	cartResponse := helper.ConvertCartModelToCreateCartResponse(updatedCartModel)
	return cartResponse, nil
}

func (cs *CartService) UpdateItem(item *dtos.Item, cart_id string) (*dtos.CreateCartResponse, error) {
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return nil, err
	}
	if cartModel.IsOrdered {
		return nil, errors.New("items ordered")
	}

	itemModel := helper.ConvertRequestItemToItemModel(item)
	itemModel.CartId = cart_id

	if err := cs.repository.UpdateItem(itemModel); err != nil {
		return nil, err
	}

	if err := cartModel.ValidateCart(); err != nil {
		return nil, err
	}

	cartResponse := helper.ConvertCartModelToCreateCartResponse(cartModel)
	return cartResponse, nil
}

func (cs *CartService) DeleteItem(item_id string, cart_id string) (*dtos.CreateCartResponse, error) {
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return nil, err
	}
	if cartModel.IsOrdered {
		return nil, errors.New("items ordered")
	}

	// Response from repository
	if err := cs.repository.DeleteItem(item_id); err != nil {
		return nil, err
	}

	if err := cartModel.ValidateCart(); err != nil {
		return nil, err
	}

	cartResponse := helper.ConvertCartModelToCreateCartResponse(cartModel)
	return cartResponse, nil
}

func (cs *CartService) CompleteOrder(cart_id string) (*dtos.CreateCartResponse, error) {
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return nil, err
	}
	if !cartModel.IsOrdered {
		cartModel.IsOrdered = true
		cartModel.OrderTime = time.Now()
		cartModel.CancelTime = time.Now().Add(5 * time.Minute)
	} else {
		return nil, errors.New("aldready ordered")
	}

	updatedCartModel, err := cs.repository.UpdateCart(cartModel)
	if err != nil {
		return nil, err
	}

	if err := updatedCartModel.ValidateCart(); err != nil {
		return nil, err
	}
	cartResponse := helper.ConvertCartModelToCreateCartResponse(updatedCartModel)
	return cartResponse, nil
}

func (cs *CartService) CancelOrder(cart_id string) error {
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
