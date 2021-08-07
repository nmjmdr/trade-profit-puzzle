package service

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"tradealgo/pkg/maxdiff"
	"tradealgo/pkg/types"
)

const NumberOfColumns = 2

var MissingColumnsInCSVErr = errors.New("Missing columns in csv")
var InvalidDataFormatErr = errors.New("Invalid data format")
var UnableToReadCSVErr = errors.New("Unable to read csv")

func parseRecord(record []string) (*types.Point, error) {
	if len(record) < NumberOfColumns {
		return nil, MissingColumnsInCSVErr
	}

	price, err := strconv.ParseFloat(record[0], 64)
	if err != nil {
		return nil, InvalidDataFormatErr
	}

	ticks, err := strconv.ParseInt(record[1], 10, 64)
	if err != nil {
		return nil, InvalidDataFormatErr
	}

	return &types.Point{
		Price: types.Cents(price),
		Ticks: types.Ticks(ticks),
	}, nil
}

func MaxProfit(reader *csv.Reader, compute func() maxdiff.Hooks) (*types.Trade, error) {
	hooks := compute()
	for {
		record, err := reader.Read()
		if err != nil && err != io.EOF {
			return nil, UnableToReadCSVErr
		}
		if err == io.EOF {
			trade, err := hooks.End()
			if err != nil {
				return nil, err
			}
			return trade, nil
		}
		point, err := parseRecord(record)
		if err != nil {
			return nil, err
		}
		err = hooks.DataPoint(point)
		if err != nil {
			return nil, err
		}
	}
}
