package profitcomputer

import (
	"math/rand"
	"tradealgo/pkg/types"
)

func increasingPoints(n int) []types.PricePoint {
	end := n + 2 // need atleast two data points
	points := make([]types.PricePoint, end)
	for i := 0; i < end; i++ {
		price, _ := types.NewPrice(types.Cents(i))
		points[i] = types.PricePoint{Price: price, Ticks: types.Ticks(i)}
	}
	return points
}

func decreasingPoints(n int) []types.PricePoint {
	end := n + 2 // need atleast two data points
	points := make([]types.PricePoint, end)
	for i := 0; i < end; i++ {
		price, _ := types.NewPrice(types.Cents(n - i))
		points[i] = types.PricePoint{Price: price, Ticks: types.Ticks(i)}
	}
	return points
}

func randomizedPoints(n int) []types.PricePoint {
	end := n + 2 // need atleast two data points
	points := make([]types.PricePoint, end)
	min := 0
	max := n
	for i := 0; i < end; i++ {
		price, _ := types.NewPrice(types.Cents((rand.Intn(max-min+1) + min)))
		points[i] = types.PricePoint{Price: price, Ticks: types.Ticks(i)}
	}
	return points
}
