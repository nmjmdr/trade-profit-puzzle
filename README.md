# Solution to find the maximun possible profit given a series of prices of a stock

## Problem statement
Suppose we could access yesterday’s stock prices as a list, where:

. The indices are the time in minutes past trade opening time, which was 10:00am local time.

. The values are the price in dollars of the Latitude Financial stock at that time.

. So if the stock cost $5 at 11:00am, stock_prices_yesterday[60] = 5.

Write an efficient function that takes an array of stock prices and returns the best profit I could have made from 1 purchase and 1 sale of 1 stock

## Compling and running the app

### Clone the project
Clone the project to a local folder
`git clone https://github.com/nmjmdr/trade-profit-puzzle.git`

### Build
Once the clone is completed, the build can be done by:

`make build`

Builds the app. This should build an executable called `app`

### Running Unit Tests
`make test` 
Runs the unit tests

### Running Benchmark
`make bench`
Runs the unit tests and the benchmark tests

### Running the application
The app can be run using the following command:
`app -file <input-path-to-csv-file> [-stream]`

`-file` option should be used to specify the input path to the CSV file

`-stream` option is optional. 
If the option `-stream` is provided then the file is be read record by record. This should be particularly useful for handling large files. 
In the `-stream` option, the file is read record by record and no allocation is done for an array. The algorithm examines the values as each record is read to determine the maximum possibe profit.

If the option is not provided, then the contents of the file are read as an array, and then the algorithm is runon the array to compute the maximum possible profit.

### Format of the input CSV
The CSV file should be provided in the following format:
```
<price in cents>,<ticks>
<price in cents>,<ticks>
<price in cents>,<ticks>
```
Example:
````
100,1257894000000000000
101,1257894060000000000
102,1257894120000000000
103,1257894180000000000
104,1257894240000000000
````
The first column represents the price in cents and the second column the price at particular time unit (represented using ticks)

An example of the CSV file is provided in `cmd/app/sample.csv`

Running the sample file:
`app -file ./cmd/app/sample.csv`

Using the stream option:
`app -file ./cmd/app/sample.csv -stream`

Sample output:
```
./app -file "../../../../generate-sample/large-sample.csv" -stream
Maximum possible profit:
        Transacton is at profit of: 9999999 cents
        Transaction details:
                Buy at price 100 cents and ticks 1
                Sell at price 10000099 cents and ticks 10000000

Using stream approach
        Time take to process: 2738427969 nanoseconcs
```

## Algorithm
The problem can be reduced to the problem of finding maximum difference between two elements of an array. Given the following constraints:
1. The smaller element (if possible) occurs before the larger element
2. The smaller element cannot be same as the larger element

One way to solve the problem would be compute the differences between every pair of elements of the array and then select the maximum difference. This would be O(n^2).

#### Algorithm with O(n) complexity

One way to solve the problem would be to:

1. Keep track of the maximum difference encountered until now
2. compute the difference between the minimum element encountered until now and the current element. 
3. If the difference encountered in step 2 is greater than the maximum difference being tracked, then difference computed in step 2 becomes the new maximum difference

Note that step 2 also involves keeping track of the minimum element encountered until now.

The code for the algorithm described above is given below:

```
func maxDiff(points []int) (int, error) {
	if len(points) < 2 {
		return 0, errors.New("Insufficient number of records")
	}

	maxDiff := points[1] - points[0]
	min := points[0]

	for i := 1; i < len(points); i++ {
		if points[i]-min > maxDiff {
			maxDiff = points[i] - min
		}
		if points[i] < min {
			min = points[i]
		}
	}
	return maxDiff, nil
}
GO-Playground link: https://play.golang.org/p/zn7oMci0hEf
```

The algorithm can be changed slightly to keep the track of the index of the minimum element and the index of the maximum element

# The problem of passing the input prices as array
One problem with the above algorithm is that it needs the prices as an array. If the input is being read from a 
large file containing millions of records, it would require the allocation of a large array.

This can be avoided by following an approach where the input file is read, parsed and processed record by record. The algorithm can be changed to support this approach.

The solution provides both approaches. As discussed earlier, the approach process the file record by record can be invoked by using the `-stream` approach.


## Benchmark
Streaming approach:
```
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^(Benchmark_StreamCalculator)$ tradealgo/service

goos: darwin
goarch: amd64
pkg: tradealgo/service
Benchmark_StreamCalculator-4   	   10000	   1242530 ns/op	  320352 B/op	   15018 allocs/op
```
Array approach:
```
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^(Benchmark_ArrayCalculator)$ tradealgo/service

goos: darwin
goarch: amd64
pkg: tradealgo/service
Benchmark_ArrayCalculator-4   	   10000	   1186871 ns/op	  579804 B/op	   10024 allocs/op
```

The following output was produced by running the application against 10 million records (file size 157 MB):
Array approach:
```
go run main.go -file "../../../../generate-sample/large-sample.csv"
Maximum possible profit:
        Transacton is at profit of: 9999999 cents
        Transaction details:
                Buy at price 100 cents and ticks 1
                Sell at price 10000099 cents and ticks 10000000

Using array approach
        Time take to process: 2944384386 nanoseconcs
```
Stream approach:
```
Maximum possible profit:
        Transacton is at profit of: 9999999 cents
        Transaction details:
                Buy at price 100 cents and ticks 1
                Sell at price 10000099 cents and ticks 10000000

Using stream approach
        Time take to process: 2633286822 nanoseconcs
```

The stream approach takes:
2944384386 - 2633286822 = 311,097,564 nanoseconds = 311 milliseconds less

_The stream approach took 311 milliseconds less than the array approach to process the file of 10 million records_