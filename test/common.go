package test

import (
	"math/rand"
	"tradealgo/pkg/types"
)

func increasingPoints(n int) []*types.Point {
	points := make([]*types.Point, n)
	for i := 0; i < n; i++ {
		points[i] = &types.Point{Price: types.Cents(i), Ticks: types.Ticks(i)}
	}
	return points
}

func decreasingPoints(n int) []*types.Point {
	points := make([]*types.Point, n)
	for i := 0; i < n; i++ {
		points[i] = &types.Point{Price: types.Cents(n - i), Ticks: types.Ticks(i)}
	}
	return points
}

func randomizedPoints(n int) []*types.Point {
	points := make([]*types.Point, n)
	min := 0
	max := n
	for i := 0; i < n; i++ {
		points[i] = &types.Point{Price: types.Cents((rand.Intn(max-min+1) + min)), Ticks: types.Ticks(i)}
	}
	return points
}
