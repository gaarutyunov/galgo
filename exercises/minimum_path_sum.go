package exercises

// MinimumPathSumV1 https://leetcode.com/problems/minimum-path-sum/
func MinimumPathSumV1(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i+j == 0 {
				continue
			}

			if i == 0 {
				dp[i][j] += grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] += grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] += min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
			}
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func MinimumPathSumV2(grid [][]int) int {
	dp := make([]int, len(grid[0]))
	dp[0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		nextRow := make([]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			if j == 0 {
				nextRow[j] += dp[j] + grid[i][j]
			} else {
				nextRow[j] += min(dp[j]+grid[i][j], nextRow[j-1]+grid[i][j])
			}
		}
		dp = nextRow
	}

	return dp[len(grid[0])-1]
}
