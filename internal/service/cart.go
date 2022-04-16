package service

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	repo "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/repository"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
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
	err := cs.repository.DeleteCart(cart_id)
	if err != nil {
		return err
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
	itemModel := helper.ConvertRequestItemToItemModel(item)
	itemModel.CartId = cart_id
	_, err := cs.repository.CreateItem(itemModel)
	if err != nil {
		return nil, err
	}
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return nil, err
	}

	if err := cartModel.ValidateCart(); err != nil {
		return nil, err
	}

	cartResponse := helper.ConvertCartModelToCreateCartResponse(cartModel)
	return cartResponse, nil
}

func (cs *CartService) UpdateItem(item *dtos.Item, cart_id string) (*dtos.CreateCartResponse, error) {
	itemModel := helper.ConvertRequestItemToItemModel(item)
	itemModel.CartId = cart_id

	if err := cs.repository.UpdateItem(itemModel); err != nil {
		return nil, err
	}

	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return nil, err
	}

	if err := cartModel.ValidateCart(); err != nil {
		return nil, err
	}

	cartResponse := helper.ConvertCartModelToCreateCartResponse(cartModel)
	return cartResponse, nil
}

func (cs *CartService) DeleteItem(item_id string, cart_id string) (*dtos.CreateCartResponse, error) {
	// Response from repository
	err := cs.repository.DeleteItem(item_id)
	if err != nil {
		return nil, err
	}
	cartModel, err := cs.repository.FetchCart(cart_id)
	if err != nil {
		return nil, err
	}

	if err := cartModel.ValidateCart(); err != nil {
		return nil, err
	}

	cartResponse := helper.ConvertCartModelToCreateCartResponse(cartModel)
	return cartResponse, nil
}
