package main

func maxProfit(prices []int) int {
	l := len(prices)
	profit := 0
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			profit += (prices[i] - prices[i-1])
		}
	}
	return profit
}
