package profitcomputer

import (
	"testing"
	"tradealgo/pkg/types"

	"github.com/stretchr/testify/suite"
)

type StreamSuite struct {
	suite.Suite
}

func (suite *StreamSuite) Test_LessThanTwoElementReturnsAnError() {
	hooks := ComputeStream()
	price, _ := types.NewPrice(1)
	err := hooks.DataPoint(&types.PricePoint{
		Price: price,
		Ticks: types.Ticks(10000),
	})
	suite.NoError(err)
	_, err = hooks.End()
	suite.Error(err)
	suite.EqualError(err, NotEnoughDataErr.Error())
}

func (suite *StreamSuite) Test_NoElementsReturnsAnError() {
	hooks := ComputeStream()
	_, err := hooks.End()
	suite.Error(err)
	suite.EqualError(err, NotEnoughDataErr.Error())
}

func (suite *StreamSuite) Test_ReturnsCorrectMaxDifference() {
	correctDifference := types.Cents(10)
	hooks := ComputeStream()
	for i, p := range []int64{10, 8, 11, 13, 15, 12, 15, 18} {
		price, _ := types.NewPrice(types.Cents(p))
		hooks.DataPoint(&types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := hooks.End()
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(1), tr.Buy.Ticks)
	suite.Equal(types.Ticks(7), tr.Sell.Ticks)
}

func (suite *StreamSuite) Test_ReturnsCorrectMaxDifferenceTwoElements() {
	correctDifference := types.Cents(1)
	hooks := ComputeStream()
	for i, p := range []float64{10, 11} {
		price, _ := types.NewPrice(types.Cents(p))
		hooks.DataPoint(&types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := hooks.End()
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(0), tr.Buy.Ticks)
	suite.Equal(types.Ticks(1), tr.Sell.Ticks)
}

func (suite *StreamSuite) Test_ReturnsCorrectMaxDifferenceIncreasingValues() {
	correctDifference := types.Cents(5)
	hooks := ComputeStream()
	for i, p := range []float64{10, 11, 12, 13, 14, 15} {
		price, _ := types.NewPrice(types.Cents(p))
		hooks.DataPoint(&types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := hooks.End()
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(0), tr.Buy.Ticks)
	suite.Equal(types.Ticks(5), tr.Sell.Ticks)
}

func (suite *StreamSuite) Test_ReturnsCorrectMaxDifferenceDecreasingValues() {
	correctDifference := types.Cents(-1)
	hooks := ComputeStream()
	for i, p := range []float64{15, 14, 13, 12, 11, 10} {
		price, _ := types.NewPrice(types.Cents(p))
		hooks.DataPoint(&types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := hooks.End()
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(0), tr.Buy.Ticks)
	suite.Equal(types.Ticks(1), tr.Sell.Ticks)
}

func (suite *StreamSuite) Test_ReturnsCorrectMaxDifferenceOccursInTheMid() {
	correctDifference := types.Cents(3)
	hooks := ComputeStream()
	for i, p := range []float64{15, 14, 13, 16, 11, 10} {
		price, _ := types.NewPrice(types.Cents(p))
		hooks.DataPoint(&types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := hooks.End()
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(2), tr.Buy.Ticks)
	suite.Equal(types.Ticks(3), tr.Sell.Ticks)
}

func (suite *StreamSuite) Test_ReturnsCorrectMaxDifferenceOccursInTheEnd() {
	correctDifference := types.Cents(2)
	hooks := ComputeStream()
	for i, p := range []float64{15, 14, 13, 12, 8, 10} {
		price, _ := types.NewPrice(types.Cents(p))
		hooks.DataPoint(&types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := hooks.End()
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(4), tr.Buy.Ticks)
	suite.Equal(types.Ticks(5), tr.Sell.Ticks)
}

func (suite *StreamSuite) Test_ReturnsCorrectMaxDifferenceOccursInTheStart() {
	correctDifference := types.Cents(3)
	hooks := ComputeStream()
	for i, p := range []float64{15, 18, 13, 12, 8, 10} {
		price, _ := types.NewPrice(types.Cents(p))
		hooks.DataPoint(&types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := hooks.End()
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(0), tr.Buy.Ticks)
	suite.Equal(types.Ticks(1), tr.Sell.Ticks)
}

func TestStreamSuite(t *testing.T) {
	suite.Run(t, new(StreamSuite))
}

type ArraySuite struct {
	suite.Suite
}

func (suite *ArraySuite) Test_LessThanTwoElementReturnsAnError() {
	price, _ := types.NewPrice(1)
	_, err := ComputeArray([]types.PricePoint{{Price: price, Ticks: types.Ticks(100000)}})
	suite.Error(err)
	suite.EqualError(err, NotEnoughDataErr.Error())
}

func (suite *ArraySuite) Test_NoElementsReturnsAnError() {
	_, err := ComputeArray([]types.PricePoint{})
	suite.Error(err)
	suite.EqualError(err, NotEnoughDataErr.Error())
}

func (suite *ArraySuite) Test_ReturnsCorrectMaxDifference() {
	correctDifference := types.Cents(10)

	pricePoints := []types.PricePoint{}
	for i, p := range []int64{10, 8, 11, 13, 15, 12, 15, 18} {
		price, _ := types.NewPrice(types.Cents(p))
		pricePoints = append(pricePoints, types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := ComputeArray(pricePoints)
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(1), tr.Buy.Ticks)
	suite.Equal(types.Ticks(7), tr.Sell.Ticks)
}

func (suite *ArraySuite) Test_ReturnsCorrectMaxDifferenceTwoElements() {
	correctDifference := types.Cents(1)
	pricePoints := []types.PricePoint{}
	for i, p := range []float64{10, 11} {
		price, _ := types.NewPrice(types.Cents(p))
		pricePoints = append(pricePoints, types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := ComputeArray(pricePoints)
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(0), tr.Buy.Ticks)
	suite.Equal(types.Ticks(1), tr.Sell.Ticks)
}

func (suite *ArraySuite) Test_ReturnsCorrectMaxDifferenceIncreasingValues() {
	correctDifference := types.Cents(5)
	pricePoints := []types.PricePoint{}
	for i, p := range []float64{10, 11, 12, 13, 14, 15} {
		price, _ := types.NewPrice(types.Cents(p))
		pricePoints = append(pricePoints, types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := ComputeArray(pricePoints)
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(0), tr.Buy.Ticks)
	suite.Equal(types.Ticks(5), tr.Sell.Ticks)
}

func (suite *ArraySuite) Test_ReturnsCorrectMaxDifferenceDecreasingValues() {
	correctDifference := types.Cents(-1)
	pricePoints := []types.PricePoint{}
	for i, p := range []float64{15, 14, 13, 12, 11, 10} {
		price, _ := types.NewPrice(types.Cents(p))
		pricePoints = append(pricePoints, types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := ComputeArray(pricePoints)
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(0), tr.Buy.Ticks)
	suite.Equal(types.Ticks(1), tr.Sell.Ticks)
}

func (suite *ArraySuite) Test_ReturnsCorrectMaxDifferenceOccursInTheMid() {
	correctDifference := types.Cents(3)
	pricePoints := []types.PricePoint{}
	for i, p := range []float64{15, 14, 13, 16, 11, 10} {
		price, _ := types.NewPrice(types.Cents(p))
		pricePoints = append(pricePoints, types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := ComputeArray(pricePoints)
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(2), tr.Buy.Ticks)
	suite.Equal(types.Ticks(3), tr.Sell.Ticks)
}

func (suite *ArraySuite) Test_ReturnsCorrectMaxDifferenceOccursInTheEnd() {
	correctDifference := types.Cents(2)
	pricePoints := []types.PricePoint{}
	for i, p := range []float64{15, 14, 13, 12, 8, 10} {
		price, _ := types.NewPrice(types.Cents(p))
		pricePoints = append(pricePoints, types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := ComputeArray(pricePoints)
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(4), tr.Buy.Ticks)
	suite.Equal(types.Ticks(5), tr.Sell.Ticks)
}

func (suite *ArraySuite) Test_ReturnsCorrectMaxDifferenceOccursAtTheStart() {
	correctDifference := types.Cents(3)
	pricePoints := []types.PricePoint{}
	for i, p := range []float64{15, 18, 13, 12, 8, 10} {
		price, _ := types.NewPrice(types.Cents(p))
		pricePoints = append(pricePoints, types.PricePoint{
			Price: price,
			Ticks: types.Ticks(i),
		})
	}
	tr, err := ComputeArray(pricePoints)
	suite.NoError(err)
	suite.Equal(correctDifference, tr.Sell.Price.Value-tr.Buy.Price.Value)
	suite.Equal(types.Ticks(0), tr.Buy.Ticks)
	suite.Equal(types.Ticks(1), tr.Sell.Ticks)
}

func TestArraySuite(t *testing.T) {
	suite.Run(t, new(ArraySuite))
}
