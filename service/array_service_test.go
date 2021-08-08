package service

import (
	"encoding/csv"
	"strings"
	"testing"
	"tradealgo/pkg/types"

	"github.com/stretchr/testify/suite"
)

type ArraySuite struct {
	suite.Suite
}

func (suite *ArraySuite) Test_InvokesHooksAndReturnsTradeOnEnd() {

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

	a := NewArrayCalculator()
	trade, err := a.Compute(csv)

	suite.NoError(err)
	suite.EqualValues(expectedTrade, trade)
}

func (suite *ArraySuite) Test_CSVParseError() {

	records := `10,1257894000000000000
11,wrongdata
`
	csv := csv.NewReader(strings.NewReader(records))

	a := NewArrayCalculator()
	_, err := a.Compute(csv)

	suite.Error(err)
	suite.EqualError(err, InvalidDataFormatErr.Error())
}

func TestArraySuite(t *testing.T) {
	suite.Run(t, new(ArraySuite))
}
