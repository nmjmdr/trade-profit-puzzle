package test

import (
	"math/rand"
	"tradealgo/pkg/types"
)

func increasingPoints(n int) []*types.TradeRecord {
	points := make([]*types.TradeRecord, n)
	for i := 0; i < n; i++ {
		points[i] = &types.TradeRecord{Price: types.Cents(i), Ticks: types.Ticks(i)}
	}
	return points
}

func decreasingPoints(n int) []*types.TradeRecord {
	points := make([]*types.TradeRecord, n)
	for i := 0; i < n; i++ {
		points[i] = &types.TradeRecord{Price: types.Cents(n - i), Ticks: types.Ticks(i)}
	}
	return points
}

func randomizedPoints(n int) []*types.TradeRecord {
	points := make([]*types.TradeRecord, n)
	min := 0
	max := n
	for i := 0; i < n; i++ {
		points[i] = &types.TradeRecord{Price: types.Cents((rand.Intn(max-min+1) + min)), Ticks: types.Ticks(i)}
	}
	return points
}
