package types

import "errors"

type Cents int64
type Ticks int64

var InvalidValueForPriceErr = errors.New("Price cannot be negative")

type Price struct {
	Value Cents
	// can include the idea of currency later
}

func NewPrice(val Cents) (Price, error) {
	if val < 0 {
		return Price{}, InvalidValueForPriceErr
	}
	return Price{
		Value: val,
	}, nil
}

type PricePoint struct {
	Price Price
	Ticks Ticks
}

type Transaction struct {
	Buy  *PricePoint
	Sell *PricePoint
}
