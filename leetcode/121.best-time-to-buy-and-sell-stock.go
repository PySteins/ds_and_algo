package leetcode

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
// 121. 买卖股票的最佳时机

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
			continue
		}
		if prices[i]-buy > max {
			max = prices[i] - buy
		}
	}
	return max
}
