package helper

import (
	"reflect"
	"testing"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
)

func TestConvertCreateCartRequestToCartModel(usecase *testing.T) {
	var expected = models.Cart{UserID: "123456789"}
	calculated := ConvertCreateCartRequestToCartModel(cartRequest)
	if !reflect.DeepEqual(expected, *calculated) {
		usecase.Errorf("Test Fail: Calculated [%v]\tExpected [%v]\n", calculated, expected)
	}
}

func TestConvertCartModelToCreateCartResponse(usecase *testing.T) {
	// var cartModel = models.Cart{
	// 	ID:         sampleID,
	// 	UserID:     sampleUserID,
	// 	Items:      sampleItems1,
	// 	Price:      samplePrice,
	// 	IsOrdered:  sampleIsOrdered,
	// 	OrderTime:  sampleOrderTime,
	// 	CancelTime: sampleCancelTime,
	// }
	// var expected = dtos.CreateCartResponse{
	// 	ID:        &sampleID,
	// 	UserID:    &sampleUserID,
	// 	Item:      sampleItems2,
	// 	Price:     &samplePrice,
	// 	IsOrdered: &sampleIsOrdered,
	// 	OrderTime: &(*strfmt.Date)(sampleOrderTime),
	// 	CancelTime: &(*strfmt.Date)(sampleCancelTime),
	// }
	expected := dtos.CreateCartResponse{}
	cartModel := models.Cart{}
	calculated := ConvertCartModelToCreateCartResponse(&cartModel)
	if !reflect.DeepEqual(expected, *calculated) {
		usecase.Errorf("Test Fail: Calculated [%v]\tExpected [%v]\n", calculated, expected)
	}
}

// var newModel = new(dtos.CreateCartResponse)

var cartRequest = &dtos.CreateCartRequest{
	UserID: &sampleUserID,
}

var sampleUserID = "123456789"

// var sampleID = "1"
// var sampleItems1 = []models.Item{
// 	{ID: "sampleItemID1", Price: 10, Quantity: 1},
// 	{ID: "sampleItemID2", Price: 10, Quantity: 1},
// }
// var sampleItems2 = []*dtos.Item{
// 	{ID: &sampleID, Price: &samplePrice, Quantity: &sampleQuantity},
// 	{ID: &sampleID, Price: &samplePrice, Quantity: &sampleQuantity},
// }
// var sampleQuantity = uint64(1)
// var samplePrice = float64(10)
// var sampleIsOrdered = true
// var sampleOrderTime = time.Now()
// var sampleCancelTime = time.Now().Add(time.Hour * 1)
