package profitcomputer

import (
	"errors"
	"tradealgo/pkg/types"
)

var NotEnoughDataErr = errors.New("Not emough data points to compute max difference")

type Hooks struct {
	DataPoint func(*types.PricePoint) error
	End       func() (*types.Transaction, error)
}

func ComputeStream() Hooks {
	firstTwo := []*types.PricePoint{}
	var maxDiff types.Cents
	var min types.Cents
	initialised := false

	tr := &types.Transaction{}
	var minRecord *types.PricePoint

	return Hooks{
		DataPoint: func(record *types.PricePoint) error {
			if !initialised {
				firstTwo = append(firstTwo, record)
				if len(firstTwo) == 2 {
					maxDiff = firstTwo[1].Price.Value - firstTwo[0].Price.Value
					min = firstTwo[0].Price.Value
					minRecord = firstTwo[0]

					tr.Buy = firstTwo[0]
					tr.Sell = firstTwo[1]

					initialised = true
				}
			}

			if initialised {
				if (record.Price.Value - min) > maxDiff {
					maxDiff = record.Price.Value - min
					tr.Sell = record
					tr.Buy = minRecord
				}
				if record.Price.Value < min {
					min = record.Price.Value
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

func ComputeArray(records []types.PricePoint) (*types.Transaction, error) {
	if len(records) < 2 {
		return nil, NotEnoughDataErr
	}

	tr := &types.Transaction{}

	maxDiff := records[1].Price.Value - records[0].Price.Value
	min := records[0].Price.Value
	tr.Buy = &records[0]
	tr.Sell = &records[1]
	minRecord := &records[0]

	for i := 1; i < len(records); i++ {
		if records[i].Price.Value-min > maxDiff {
			maxDiff = records[i].Price.Value - min
			tr.Buy = minRecord
			tr.Sell = &records[i]
		}
		if records[i].Price.Value < min {
			min = records[i].Price.Value
			minRecord = &records[i]
		}
	}
	return tr, nil
}
