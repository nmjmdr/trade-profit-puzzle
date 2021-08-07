package types

type Cents float64
type Ticks int64

type TradeRecord struct {
	Price Cents
	Ticks Ticks
}

type Transaction struct {
	Buy    *TradeRecord
	Sell   *TradeRecord
	Profit Cents
}
