package lc

import "math"

func maxProfitBruteForce(prices []int) int {
	var profit int
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}
	return profit
}

func maxProfit(prices []int) int {
	minPrice, maxProfit := math.MaxInt, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

func maxProfitPointers(prices []int) int {
	maxprofit, buy, sell := 0, 0, 1
	for sell < len(prices) {
		if prices[buy] < prices[sell] {
			if prices[sell]-prices[buy] > maxprofit {
				maxprofit = prices[sell] - prices[buy]
			}
		} else {
			buy = sell
		}
		sell++
	}
	return maxprofit
}
