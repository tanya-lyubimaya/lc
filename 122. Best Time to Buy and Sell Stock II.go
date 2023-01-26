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
