package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
	"time"
	"tradealgo/pkg/maxdiff"
	"tradealgo/service"
)

func main() {
	var buffer bytes.Buffer

	const Max = 10000000
	const BaseTick = 100000
	const BasePrice = 100
	for index := 0; index < Max; index++ {
		buffer.WriteString(strings.Join([]string{strconv.Itoa(BasePrice + index), strconv.Itoa(BaseTick + index)}, ","))
		buffer.Write([]byte(fmt.Sprintln("")))
	}
	input := buffer.String()

	reader := csv.NewReader(strings.NewReader(input))
	start := time.Now()
	_, err := service.MaxProfit(reader, maxdiff.MaxDiffCompute)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	duration := time.Since(start)
	fmt.Printf("Time taken to process %d input records: %#v millseconds, or %v nanoseconds per record", Max, duration.Milliseconds(), duration.Nanoseconds()/Max)
}
