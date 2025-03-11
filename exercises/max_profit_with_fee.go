package exercises

func MaxProfitWithFeeV1(prices []int, fee int) int {
	free, hold := make([]int, len(prices), len(prices)), make([]int, len(prices), len(prices))
	hold[0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		hold[i] = max(hold[i-1], free[i-1]-prices[i])
		free[i] = max(free[i-1], hold[i-1]+prices[i]-fee)
	}

	return free[len(prices)-1]
}

func MaxProfitWithFeeV2(prices []int, fee int) int {
	free, hold := 0, -prices[0]

	for i := 1; i < len(prices); i++ {
		newHold := max(hold, free-prices[i])
		newFree := max(free, hold+prices[i]-fee)
		hold = newHold
		free = newFree
	}

	return free
}
