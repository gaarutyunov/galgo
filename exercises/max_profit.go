package exercises

// MaxProfitV1 https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
func MaxProfitV1(k int, prices []int) int {
	memo := make([][][]int, len(prices))
	for i := range memo {
		memo[i] = make([][]int, 2)
		for j := range memo[i] {
			memo[i][j] = make([]int, k)
			for l := range memo[i][j] {
				memo[i][j][l] = -1
			}
		}
	}

	return maxProfitDp(k, prices, 0, 0, memo)
}

func maxProfitDp(remaining int, prices []int, i int, holding int, memo [][][]int) (ans int) {
	if i == len(prices) || remaining == 0 {
		return 0
	}

	if memo[i][holding][remaining-1] != -1 {
		return memo[i][holding][remaining-1]
	}

	ans = maxProfitDp(remaining, prices, i+1, holding, memo)

	if holding == 1 {
		ans = max(
			prices[i]+maxProfitDp(remaining-1, prices, i+1, 0, memo),
			ans,
		)
	} else {
		ans = max(
			-prices[i]+maxProfitDp(remaining, prices, i+1, 1, memo),
			ans,
		)
	}

	memo[i][holding][remaining-1] = ans

	return
}

func MaxProfitV2(k int, prices []int) int {
	dp := make([][][]int, len(prices)+1)
	for i := range dp {
		dp[i] = [][]int{
			make([]int, k+1),
			make([]int, k+1),
		}
	}

	for i := len(prices) - 1; i >= 0; i-- {
		for remaining := 1; remaining <= k; remaining++ {
			for holding := 0; holding < 2; holding++ {
				ans := dp[i+1][holding][remaining]

				if holding == 1 {
					ans = max(ans, prices[i]+dp[i+1][0][remaining-1])
				} else {
					ans = max(ans, -prices[i]+dp[i+1][1][remaining])
				}

				dp[i][holding][remaining] = ans
			}
		}
	}

	return dp[0][0][k]
}

func MaxProfitV3(k int, prices []int) int {
	free := make([][]int, len(prices))
	hold := make([][]int, len(prices))
	for i := range free {
		free[i] = make([]int, k+1)
	}
	hold[0] = make([]int, k+1)
	for i := range hold[0] {
		hold[0][i] = prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for remaining := 1; remaining <= k; remaining++ {
			free[i][remaining] = max(free[i-1][remaining], prices[i]+hold[i-1][remaining-1])
			hold[i][remaining] = max(hold[i-1][remaining], -prices[i]+free[i-1][remaining])
		}
	}

	return free[len(prices)-1][k]
}
