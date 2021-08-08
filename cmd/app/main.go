package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	"tradealgo/pkg/profitcomputer"
	"tradealgo/pkg/types"
	"tradealgo/service"
)

func printUsage(message string) {
	fmt.Printf("error: %s", message)
	fmt.Println()
	fmt.Println("usage: app -file <path-to-csv> [-stream]")
	fmt.Println("CSV format: ")
	fmt.Println("<price-in-cents>,<ticks>")
}

const UnableToParseInputErrCode = 1
const UnableToReadInputFile = 2
const UnableToProcessCSV = 3

func printTransaction(transaction *types.Transaction) {
	fmt.Println("Maximum possible profit:")
	var delta = transaction.Sell.Price.Value - transaction.Buy.Price.Value
	if delta < 0 {
		fmt.Printf("\tTransacton is at loss of: %v cents", delta)
	} else {
		fmt.Printf("\tTransacton is at profit of: %v cents", delta)
	}
	fmt.Println("")
	fmt.Println("\tTransaction details:")
	fmt.Printf("\t\tBuy at price %v cents and ticks %v", transaction.Buy.Price.Value, transaction.Buy.Ticks)
	fmt.Println("")
	fmt.Printf("\t\tSell at price %v cents and ticks %v", transaction.Sell.Price.Value, transaction.Sell.Ticks)
}

func printStats(useStream bool, timeTaken time.Duration) {
	fmt.Println("")
	fmt.Println("")
	if useStream {
		fmt.Println("Using stream approach")
	} else {
		fmt.Println("Using array approach")
	}
	fmt.Printf("\tTime take to process: %v nanoseconcs", timeTaken.Nanoseconds())
	fmt.Println("")
}

func main() {
	filePathFlag := flag.String("file", "", "-file <path-to-csv>")
	useStreamFlag := flag.Bool("stream", false, "-stream")

	flag.Parse()

	filePath := *filePathFlag
	useStream := *useStreamFlag

	filePath = strings.TrimSpace(filePath)

	if len(filePath) == 0 {
		printUsage("missing input file path")
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("unable to read input file: ", err)
		os.Exit(UnableToReadInputFile)
	}

	var calc service.MaxProfitCalculator
	start := time.Now()
	if useStream {
		calc = service.NewStreamCalculator(profitcomputer.ComputeStream)

	} else {
		calc = service.NewArrayCalculator()
	}

	transaction, err := calc.Compute(csv.NewReader(file))
	if err != nil {
		fmt.Println("unable to process csv file")
		os.Exit(UnableToProcessCSV)
	}
	since := time.Since(start)

	printTransaction(transaction)
	printStats(useStream, since)
}
