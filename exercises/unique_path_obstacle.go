package exercises

// UniquePathObstacleV1 https://leetcode.com/problems/unique-paths-ii
func UniquePathObstacleV1(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	dp[0][0] = 1

	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if i+j == 0 {
				continue
			}

			if obstacleGrid[i][j] != 1 {
				if i > 0 && dp[i-1][j] > 0 {
					dp[i][j] += max(1, dp[i-1][j])
				}
				if j > 0 && dp[i][j-1] > 0 {
					dp[i][j] += max(1, dp[i][j-1])
				}
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func UniquePathObstacleV2(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	dp := make([]int, len(obstacleGrid[0]))
	for j := 0; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[j] = 1
	}

	for i := 1; i < len(obstacleGrid); i++ {
		nextRow := make([]int, len(obstacleGrid[0]))

		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] != 1 {
				if i > 0 && dp[j] > 0 {
					nextRow[j] += max(1, dp[j])
				}
				if j > 0 && nextRow[j-1] > 0 {
					nextRow[j] += max(1, nextRow[j-1])
				}
			}
		}

		dp = nextRow
	}

	return dp[len(dp)-1]
}

func UniquePathObstacleV3(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	obstacleGrid[0][0] = 1
	for j := 1; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] == 1 {
			obstacleGrid[0][j] = 0
		} else {
			obstacleGrid[0][j] = obstacleGrid[0][j-1]
		}
	}
	for i := 1; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
		} else {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		}
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}

			obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
		}
	}

	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
