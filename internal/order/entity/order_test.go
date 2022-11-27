package entity

// go mod tidy to download a package
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test given that I have (blank ID) when I create a new order then it should return error
func TestGivenAnEmptyID_WhenCreateNewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "1"}
	assert.Error(t, order.IsValid(), "invalid id")
}
func TestGivenAnPrice_WhenCreateNewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnTax_WhenCreateNewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.IsValid(), "invalid Tax")
}

func TestGivenAnValidParams_whenICallNewOrder_thenIshouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 2.0}

	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.IsValid())
}

// (*testing.T) - o T and to be able to run the tests
func TestGivenAnValidParams_whenICallNewOrderFunc_thenIshouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)

	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.IsValid())
}

func TestGivenAnPriceAndTax_whenICallCalculatePrice_thenIshouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.0, order.FinalPrice)

}
