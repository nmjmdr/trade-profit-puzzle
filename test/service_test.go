package test

import (
	"encoding/csv"
	"strings"
	"testing"
	"tradealgo/pkg/maxdiff"
	"tradealgo/pkg/types"
	"tradealgo/service"

	"github.com/stretchr/testify/assert"
)

func Test_InvokesHooksAndReturnsTradeOnEnd(t *testing.T) {

	records := `10,1257894000000000000
11,1257894000000000001
`
	csv := csv.NewReader(strings.NewReader(records))

	expectedTrade := &types.Trade{
		BuyPoint:  &types.Point{Price: types.Cents(10), Ticks: types.Ticks(1257894000000000000)},
		SellPoint: &types.Point{Price: types.Cents(11), Ticks: types.Ticks(1257894000000000001)},
	}

	pointsReceived := []*types.Point{}

	trade, err := service.MaxProfit(csv, func() maxdiff.Hooks {
		return maxdiff.Hooks{
			DataPoint: func(pt *types.Point) error {
				pointsReceived = append(pointsReceived, pt)
				return nil
			},
			End: func() (*types.Trade, error) {
				return &types.Trade{
					BuyPoint:  pointsReceived[0],
					SellPoint: pointsReceived[1],
				}, nil
			},
		}
	})

	assert.NoError(t, err)
	assert.EqualValues(t, expectedTrade, trade)
}

func Test_CSVParseError(t *testing.T) {

	records := `10,1257894000000000000
11,wrongdata
`
	csv := csv.NewReader(strings.NewReader(records))

	pointsReceived := []*types.Point{}

	_, err := service.MaxProfit(csv, func() maxdiff.Hooks {
		return maxdiff.Hooks{
			DataPoint: func(pt *types.Point) error {
				pointsReceived = append(pointsReceived, pt)
				return nil
			},
			End: func() (*types.Trade, error) {
				return &types.Trade{
					BuyPoint:  pointsReceived[0],
					SellPoint: pointsReceived[1],
				}, nil
			},
		}
	})

	assert.Error(t, err)
	assert.EqualError(t, err, service.InvalidDataFormatErr.Error())
}
