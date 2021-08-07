package maxdiff

import (
	"errors"
	"tradealgo/pkg/types"
)

var NotEnoughDataErr = errors.New("Not emough data points to compute max difference")

type Hooks struct {
	DataPoint func(*types.Point) error
	End       func() (*types.Trade, error)
}

func MaxDiffCompute() Hooks {
	firstTwo := []*types.Point{}
	var maxDiff types.Cents
	var min types.Cents
	initialised := false

	trade := &types.Trade{}
	var minPoint *types.Point

	return Hooks{
		DataPoint: func(point *types.Point) error {
			if !initialised {
				firstTwo = append(firstTwo, point)
				if len(firstTwo) == 2 {
					maxDiff = firstTwo[1].Price - firstTwo[0].Price
					min = firstTwo[0].Price

					trade.BuyPoint = firstTwo[0]
					trade.SellPoint = firstTwo[1]
					trade.Delta = maxDiff

					initialised = true
				}
			}

			if initialised {
				if (point.Price - min) > maxDiff {
					maxDiff = point.Price - min

					trade.SellPoint = point
					trade.BuyPoint = minPoint
					trade.Delta = maxDiff
				}
				if point.Price < min {
					min = point.Price
					minPoint = point
				}
			}
			return nil
		},
		End: func() (*types.Trade, error) {
			if !initialised {
				return nil, NotEnoughDataErr
			}
			return trade, nil
		},
	}
}

// the above version which works on a record by record is derived
// from the below code I came up with which works on an array as input:

/*
type Result struct {
	Delta      int
	StartIndex int
	EndIndex   int
}

func maxDiff(points []int) (*Result, error) {
	if len(points) < 2 {
		return nil, errors.New("Insufficient number of records")
	}

	maxDiff := points[1] - points[0]
	min := points[0]
	result := &Result{StartIndex: 0, EndIndex: 1}
	currentMinIndex := 0

	for i := 1; i < len(points); i++ {
		if points[i]-min > maxDiff {
			maxDiff = points[i] - min
			result.StartIndex = currentMinIndex
			result.EndIndex = i
			result.Delta = maxDiff
		}
		if points[i] < min {
			min = points[i]
			currentMinIndex = i
		}
	}
	return result, nil
}
*/
