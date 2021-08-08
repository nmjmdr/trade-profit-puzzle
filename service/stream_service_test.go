package service

import (
	"encoding/csv"
	"strings"
	"testing"
	"tradealgo/pkg/profitcomputer"
	"tradealgo/pkg/types"

	"github.com/stretchr/testify/suite"
)

type StreamSuite struct {
	suite.Suite
}

func (suite *StreamSuite) Test_InvokesHooksAndReturnsTradeOnEnd() {

	records := `10,1257894000000000000
11,1257894000000000001
`
	csv := csv.NewReader(strings.NewReader(records))

	buyPrice, _ := types.NewPrice(types.Cents(10))
	sellPrice, _ := types.NewPrice(types.Cents(11))
	expectedTrade := &types.Transaction{
		Buy:  &types.PricePoint{Price: buyPrice, Ticks: types.Ticks(1257894000000000000)},
		Sell: &types.PricePoint{Price: sellPrice, Ticks: types.Ticks(1257894000000000001)},
	}

	pointsReceived := []*types.PricePoint{}

	sb := NewStreamCalculator(func() profitcomputer.Hooks {
		return profitcomputer.Hooks{
			DataPoint: func(pt *types.PricePoint) error {
				pointsReceived = append(pointsReceived, pt)
				return nil
			},
			End: func() (*types.Transaction, error) {
				return &types.Transaction{
					Buy:  pointsReceived[0],
					Sell: pointsReceived[1],
				}, nil
			},
		}
	})
	trade, err := sb.Compute(csv)

	suite.NoError(err)
	suite.EqualValues(expectedTrade, trade)
}

func (suite *StreamSuite) Test_CSVParseError() {

	records := `10,1257894000000000000
11,wrongdata
`
	csv := csv.NewReader(strings.NewReader(records))

	pointsReceived := []*types.PricePoint{}

	sb := NewStreamCalculator(func() profitcomputer.Hooks {
		return profitcomputer.Hooks{
			DataPoint: func(pt *types.PricePoint) error {
				pointsReceived = append(pointsReceived, pt)
				return nil
			},
			End: func() (*types.Transaction, error) {
				return &types.Transaction{
					Buy:  pointsReceived[0],
					Sell: pointsReceived[1],
				}, nil
			},
		}
	})
	_, err := sb.Compute(csv)

	suite.Error(err)
	suite.EqualError(err, InvalidDataFormatErr.Error())
}

func TestStreamSuite(t *testing.T) {
	suite.Run(t, new(StreamSuite))
}
