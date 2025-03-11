package exercises

// UniquePathsV1 https://leetcode.com/problems/unique-paths
func UniquePathsV1(m, n int) int {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return uniquePathsDp(m, n, m-1, n-1, memo)
}

func uniquePathsDp(m, n, row, col int, memo [][]int) (ans int) {
	if row+col == 0 {
		return 1
	}

	if memo[row][col] != -1 {
		return memo[row][col]
	}

	if row > 0 {
		ans += uniquePathsDp(m, n, row-1, col, memo)
	}

	if col > 0 {
		ans += uniquePathsDp(m, n, row, col-1, memo)
	}

	memo[row][col] = ans

	return
}

func UniquePathsV2(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i+j == 0 {
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func UniquePathsV3(m, n int) int {
	dp := make([]int, n)
	dp[0] = 1

	for i := 0; i < m; i++ {
		nextRow := make([]int, n)
		for j := 0; j < n; j++ {
			if i+j == 0 {
				continue
			}

			nextRow[j] += dp[j]
			if j > 0 {
				nextRow[j] += nextRow[j-1]
			}
		}
		dp = nextRow
	}

	return dp[n-1]
}
