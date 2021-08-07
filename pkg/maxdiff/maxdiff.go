package maxdiff

import (
	"errors"
	"tradealgo/pkg/types"
)

var NotEnoughDataErr = errors.New("Not emough data points to compute max difference")

type Hooks struct {
	DataPoint func(*types.TradeRecord) error
	End       func() (*types.Transaction, error)
}

func MaxDiffCompute() Hooks {
	firstTwo := []*types.TradeRecord{}
	var maxDiff types.Cents
	var min types.Cents
	initialised := false

	tr := &types.Transaction{}
	var minRecord *types.TradeRecord

	return Hooks{
		DataPoint: func(record *types.TradeRecord) error {
			if !initialised {
				firstTwo = append(firstTwo, record)
				if len(firstTwo) == 2 {
					maxDiff = firstTwo[1].Price - firstTwo[0].Price
					min = firstTwo[0].Price

					tr.Buy = firstTwo[0]
					tr.Sell = firstTwo[1]
					tr.Profit = maxDiff

					initialised = true
				}
			}

			if initialised {
				if (record.Price - min) > maxDiff {
					maxDiff = record.Price - min

					tr.Sell = record
					tr.Buy = minRecord
					tr.Profit = maxDiff
				}
				if record.Price < min {
					min = record.Price
					minRecord = record
				}
			}
			return nil
		},
		End: func() (*types.Transaction, error) {
			if !initialised {
				return nil, NotEnoughDataErr
			}
			return tr, nil
		},
	}
}
