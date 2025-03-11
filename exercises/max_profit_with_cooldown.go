package exercises

// MaxProfitWithCooldownV1 https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown
func MaxProfitWithCooldownV1(prices []int) int {
	hold, sold, reset := make([]int, len(prices)), make([]int, len(prices)), make([]int, len(prices))
	hold[0] = -prices[0]

	for i := 0; i < len(prices); i++ {
		sold[i] = hold[i-1] + prices[i]
		hold[i] = max(hold[i-1], reset[i-1]-prices[i])
		reset[i] = max(reset[i-1], sold[i-1])
	}

	return max(sold[len(sold)-1], reset[len(sold)-1])
}

func MaxProfitWithCooldownV2(prices []int) int {
	hold, sold, reset := -prices[0], 0, 0

	for i := 0; i < len(prices); i++ {
		newSold := hold + prices[i]
		newHold := max(hold, reset-prices[i])
		reset = max(reset, sold)
		sold = newSold
		hold = newHold
	}

	return max(sold, reset)
}
