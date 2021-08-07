package test

import (
	"math/rand"
	"testing"
	"time"
	"tradealgo/pkg/maxdiff"
	"tradealgo/pkg/types"

	"github.com/stretchr/testify/assert"
)

func Test_LessThanTwoElementReturnsAnError(t *testing.T) {
	hooks := maxdiff.MaxDiffCompute()
	err := hooks.DataPoint(&types.Point{
		Price: types.Cents(1),
		Ticks: types.Ticks(10000),
	})
	assert.NoError(t, err)
	_, err = hooks.End()
	assert.Error(t, err)
	assert.EqualError(t, err, maxdiff.NotEnoughDataErr.Error())
}

func Test_NoElementsReturnsAnError(t *testing.T) {
	hooks := maxdiff.MaxDiffCompute()
	_, err := hooks.End()
	assert.Error(t, err)
	assert.EqualError(t, err, maxdiff.NotEnoughDataErr.Error())
}

func Test_ReturnsCorrectMaxDifference(t *testing.T) {
	correctDifference := types.Cents(10)
	hooks := maxdiff.MaxDiffCompute()
	for i, p := range []float64{10, 8, 11, 13, 15, 12, 15, 18} {
		hooks.DataPoint(&types.Point{
			Price: types.Cents(p),
			Ticks: types.Ticks(i),
		})
	}
	d, err := hooks.End()
	assert.NoError(t, err)
	assert.Equal(t, correctDifference, d.Delta)
	assert.Equal(t, types.Ticks(1), d.BuyPoint.Ticks)
	assert.Equal(t, types.Ticks(7), d.SellPoint.Ticks)
}

func Test_ReturnsCorrectMaxDifferenceTwoElements(t *testing.T) {
	correctDifference := types.Cents(1)
	hooks := maxdiff.MaxDiffCompute()
	for i, p := range []float64{10, 11} {
		hooks.DataPoint(&types.Point{
			Price: types.Cents(p),
			Ticks: types.Ticks(i),
		})
	}
	d, err := hooks.End()
	assert.NoError(t, err)
	assert.Equal(t, correctDifference, d.Delta)
	assert.Equal(t, types.Ticks(0), d.BuyPoint.Ticks)
	assert.Equal(t, types.Ticks(1), d.SellPoint.Ticks)
}

func Test_GivenDecreasingElementsReturnsNegativeDifference(t *testing.T) {
	correctDifference := types.Cents(-1)
	hooks := maxdiff.MaxDiffCompute()
	for i, p := range []float64{6, 5, 4} {
		hooks.DataPoint(&types.Point{
			Price: types.Cents(p),
			Ticks: types.Ticks(i),
		})
	}
	d, err := hooks.End()
	assert.NoError(t, err)
	assert.Equal(t, correctDifference, d.Delta)
	assert.Equal(t, types.Ticks(0), d.BuyPoint.Ticks)
	assert.Equal(t, types.Ticks(1), d.SellPoint.Ticks)
}

func Test_HandlesNegativeNumbers(t *testing.T) {
	correctDifference := types.Cents(-1)
	hooks := maxdiff.MaxDiffCompute()
	for i, p := range []float64{-1, -2, -3} {
		hooks.DataPoint(&types.Point{
			Price: types.Cents(p),
			Ticks: types.Ticks(i),
		})
	}
	d, err := hooks.End()
	assert.NoError(t, err)
	assert.Equal(t, correctDifference, d.Delta)
	assert.Equal(t, types.Ticks(0), d.BuyPoint.Ticks)
	assert.Equal(t, types.Ticks(1), d.SellPoint.Ticks)
}

func Test_HandlesFirstPositiveAndThenNegativeNumbers(t *testing.T) {
	correctDifference := types.Cents(-3)
	hooks := maxdiff.MaxDiffCompute()
	for i, p := range []float64{1, -2} {
		hooks.DataPoint(&types.Point{
			Price: types.Cents(p),
			Ticks: types.Ticks(i),
		})
	}
	d, err := hooks.End()
	assert.NoError(t, err)
	assert.Equal(t, correctDifference, d.Delta)
	assert.Equal(t, types.Ticks(0), d.BuyPoint.Ticks)
	assert.Equal(t, types.Ticks(1), d.SellPoint.Ticks)
}

func Test_HandlesNegativeIncreasingNumbers(t *testing.T) {
	correctDifference := types.Cents(-1)
	hooks := maxdiff.MaxDiffCompute()
	for i, p := range []float64{-1, -2, -3} {
		hooks.DataPoint(&types.Point{
			Price: types.Cents(p),
			Ticks: types.Ticks(i),
		})
	}
	d, err := hooks.End()
	assert.NoError(t, err)
	assert.Equal(t, correctDifference, d.Delta)
	assert.Equal(t, types.Ticks(0), d.BuyPoint.Ticks)
	assert.Equal(t, types.Ticks(1), d.SellPoint.Ticks)
}

func Benchmark_IncreasingPoints(b *testing.B) {
	points := increasingPoints(b.N)
	hooks := maxdiff.MaxDiffCompute()
	for n := 0; n < b.N; n++ {
		for _, pt := range points {
			hooks.DataPoint(pt)
		}
		hooks.End()
	}
}

func Benchmark_DecreasingPoints(b *testing.B) {
	points := decreasingPoints(b.N)
	hooks := maxdiff.MaxDiffCompute()
	for n := 0; n < b.N; n++ {
		for _, pt := range points {
			hooks.DataPoint(pt)
		}
		hooks.End()
	}
}

func Benchmark_RandomizedPoints(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	points := randomizedPoints(b.N)
	hooks := maxdiff.MaxDiffCompute()
	for n := 0; n < b.N; n++ {
		for _, pt := range points {
			hooks.DataPoint(pt)
		}
		hooks.End()
	}
}
