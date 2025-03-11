package exercises

import "math"

// CoinChangeV1 https://leetcode.com/problems/coin-change/
func CoinChangeV1(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = math.MaxInt
	}

	return coinChangeDp(coins, amount, memo)
}

func coinChangeDp(coins []int, remaining int, memo []int) int {
	if remaining < 0 {
		return -1
	}

	if remaining == 0 {
		return 0
	}

	if memo[remaining] != math.MaxInt {
		return memo[remaining]
	}

	res := -1

	for _, coin := range coins {
		v := coinChangeDp(coins, remaining-coin, memo)

		if v == -1 {
			continue
		}

		if res == -1 || v < res {
			res = v + 1
		}
	}

	memo[remaining] = res

	return res
}

func CoinChangeV2(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1

		for _, coin := range coins {
			if coin > i {
				continue
			}

			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
