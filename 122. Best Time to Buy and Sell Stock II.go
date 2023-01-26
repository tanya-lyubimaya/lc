package lc

func maxProfit(prices []int) int {
	maxprofit, buy, sell := 0, 0, 1
	for sell < len(prices) {
		if prices[buy] < prices[sell] {
			if prices[sell]-prices[buy] > 0 {
				maxprofit += prices[sell] - prices[buy]
			}
		}
		buy = sell
		sell++
	}
	return maxprofit
}

func maxProfitBruteForce(prices []int) int {
	res := calculate(prices, 0)
	return res
}

func calculate(prices []int, s int) int {
	if s >= len(prices) {
		return 0
	}
	var max int
	for i := s; i < len(prices); i++ {
		var maxprofit int
		for j := i + 1; j < len(prices); j++ {
			profit := calculate(prices, j+1) + prices[j] - prices[i]
			if profit > maxprofit {
				maxprofit = profit
			}
		}
		if maxprofit > max {
			max = maxprofit
		}
	}
	return max
}
