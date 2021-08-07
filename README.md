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
