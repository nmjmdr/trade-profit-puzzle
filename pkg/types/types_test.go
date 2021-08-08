package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PriceSetsValue(t *testing.T) {
	v := Cents(10)
	p, err := NewPrice(v)
	assert.NoError(t, err)
	assert.Equal(t, p.Value, v)
}

func Test_PriceReturnsErrorOnNegativeValueInput(t *testing.T) {
	v := Cents(-10)
	_, err := NewPrice(v)
	assert.Error(t, err)
	assert.EqualError(t, err, InvalidValueForPriceErr.Error())
}
