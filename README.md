# Solution to find the maximun possible profit given a series of prices of a stock

## Problem statement
Suppose we could access yesterday’s stock prices as a list, where:

. The indices are the time in minutes past trade opening time, which was 10:00am local time.

. The values are the price in dollars of the Latitude Financial stock at that time.

. So if the stock cost $5 at 11:00am, stock_prices_yesterday[60] = 5.

Write an efficient function that takes an array of stock prices and returns the best profit I could have made from 1 purchase and 1 sale of 1 stock

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

This would require the entire 