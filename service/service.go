package service

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"tradealgo/pkg/profitcomputer"
	"tradealgo/pkg/types"
)

const NumberOfColumns = 2

var MissingColumnsInCSVErr = errors.New("Missing columns in csv")
var InvalidDataFormatErr = errors.New("Invalid data format")
var UnableToReadCSVErr = errors.New("Unable to read csv")

func parseRecord(record []string) (types.PricePoint, error) {
	if len(record) < NumberOfColumns {
		return types.PricePoint{}, MissingColumnsInCSVErr
	}
	priceIntValue, err := strconv.ParseInt(record[0], 10, 64)
	if err != nil {
		return types.PricePoint{}, InvalidDataFormatErr
	}
	price, err := types.NewPrice(types.Cents(priceIntValue))
	if err != nil {
		return types.PricePoint{}, err
	}

	ticks, err := strconv.ParseInt(record[1], 10, 64)
	if err != nil {
		return types.PricePoint{}, InvalidDataFormatErr
	}

	return types.PricePoint{
		Price: price,
		Ticks: types.Ticks(ticks),
	}, nil
}

type MaxProfitCalculator interface {
	Compute(reader *csv.Reader) (*types.Transaction, error)
}

type streamCalc struct {
	computeFn func() profitcomputer.Hooks
}

func NewStreamCalculator(compute func() profitcomputer.Hooks) MaxProfitCalculator {
	return &streamCalc{
		computeFn: compute,
	}
}

func (s *streamCalc) Compute(reader *csv.Reader) (*types.Transaction, error) {
	hooks := s.computeFn()
	for {
		record, err := reader.Read()
		if err != nil && err != io.EOF {
			return nil, UnableToReadCSVErr
		}
		if err == io.EOF {
			tr, err := hooks.End()
			if err != nil {
				return nil, err
			}
			return tr, nil
		}
		pr, err := parseRecord(record)
		if err != nil {
			return nil, err
		}
		err = hooks.DataPoint(&pr)
		if err != nil {
			return nil, err
		}
	}
}

type arrayCalc struct {
}

func NewArrayCalculator() MaxProfitCalculator {
	return &arrayCalc{}
}

func (a *arrayCalc) Compute(reader *csv.Reader) (*types.Transaction, error) {
	pricePoints := []types.PricePoint{}
	for {
		record, err := reader.Read()
		if err != nil && err != io.EOF {
			return nil, UnableToReadCSVErr
		}
		if err == io.EOF {
			break
		}
		pr, err := parseRecord(record)
		if err != nil {
			return nil, err
		}
		pricePoints = append(pricePoints, pr)
	}
	tr, err := profitcomputer.ComputeArray(pricePoints)
	return tr, err
}
