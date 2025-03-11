package exercises

func MaxValueOfCoinsV1(piles [][]int, k int) int {
	memo := make([][]int, len(piles), len(piles))
	for i := range memo {
		memo[i] = make([]int, k+1, k+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return maxValueOfCoinsDp(piles, 0, k, memo)
}

func maxValueOfCoinsDp(piles [][]int, i, k int, memo [][]int) (ans int) {
	if i == len(piles) || k == 0 {
		return 0
	}

	if memo[i][k] != -1 {
		return memo[i][k]
	}

	ans = maxValueOfCoinsDp(piles, i+1, k, memo)
	var curr int

	for j := 0; j < min(k, len(piles[i])); j++ {
		curr += piles[i][j]
		ans = max(ans, curr+maxValueOfCoinsDp(piles, i+1, k-j-1, memo))
	}

	memo[i][k] = ans

	return ans
}

func MaxValueOfCoinsV2(piles [][]int, k int) int {
	dp := make([][]int, len(piles)+1, len(piles)+1)
	for i := range dp {
		dp[i] = make([]int, k+1, k+1)
	}

	for i := len(piles) - 1; i >= 0; i-- {
		var curr int
		for remain := 0; remain < k; remain++ {
			dp[i][remain] = dp[i+1][remain]

			for j := 0; j < min(remain, len(piles[i])); j++ {
				curr += piles[i][j]
				dp[i][remain] = max(dp[i][remain], curr+dp[i+1][k-j-1])
			}
		}
	}

	return dp[0][k]
}
