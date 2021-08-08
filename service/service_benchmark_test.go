package service

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"tradealgo/pkg/profitcomputer"
)

func generateCSVReader(n int) *csv.Reader {
	var buffer bytes.Buffer

	const BaseTick = 100000
	const BasePrice = 100
	end := n + 2 // generate two more than input, as we need atleast 2 records
	for index := 0; index < end; index++ {
		buffer.WriteString(strings.Join([]string{strconv.Itoa(BasePrice + index), strconv.Itoa(BaseTick + index)}, ","))
		buffer.Write([]byte(fmt.Sprintln("")))
	}
	input := buffer.String()
	return csv.NewReader(strings.NewReader(input))
}

func Benchmark_StreamCalculator(b *testing.B) {
	sb := NewStreamCalculator(profitcomputer.ComputeStream)
	for i := 1; i < b.N; i++ {
		b.StopTimer()
		reader := generateCSVReader(i)
		b.StartTimer()
		_, err := sb.Compute(reader)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_ArrayCalculator(b *testing.B) {
	a := NewArrayCalculator()
	for i := 1; i < b.N; i++ {
		b.StopTimer()
		reader := generateCSVReader(i)
		b.StartTimer()
		_, err := a.Compute(reader)
		if err != nil {
			b.Error(err)
		}
	}
}
