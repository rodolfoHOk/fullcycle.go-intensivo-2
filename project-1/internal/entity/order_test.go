package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_if_it_gets_an_error_if_id_is_blank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}

func Test_if_it_gets_an_error_if_price_is_blank(t *testing.T) {
	order := Order{
		ID: "123",
	}
	assert.Error(t, order.Validate(), "invalid price")
}

func Test_if_it_gets_an_error_if_tax_is_blank(t *testing.T) {
	order := Order{
		ID: "123",
		Price: 10.0,
	}
	assert.Error(t, order.Validate(), "invalid tax")
}

func Test_with_all_valid_params(t *testing.T) {
	order := Order{
		ID: "123", 
		Price: 10.0, 
		Tax: 1.0,
	}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
