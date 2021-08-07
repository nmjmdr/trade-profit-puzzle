package types

type Cents float64
type Ticks int64

type Point struct {
	Price Cents
	Ticks Ticks
}

type Trade struct {
	BuyPoint  *Point
	SellPoint *Point
	Delta     Cents
}
