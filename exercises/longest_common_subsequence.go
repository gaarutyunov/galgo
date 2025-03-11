package exercises

// LongestCommonSubsequenceV1 https://leetcode.com/problems/longest-common-subsequence
func LongestCommonSubsequenceV1(text1, text2 string) int {
	memo := make([][]int, len(text1))
	for i := range memo {
		memo[i] = make([]int, len(text2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return longestCommonSubsequenceDp(text1, text2, 0, 0, memo)
}

func longestCommonSubsequenceDp(text1, text2 string, i, j int, memo [][]int) int {
	if i == len(text1) || j == len(text2) {
		return 0
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	var ans int

	if text1[i] == text2[j] {
		ans = 1 + longestCommonSubsequenceDp(text1, text2, i+1, j+1, memo)
	} else {
		ans = max(
			longestCommonSubsequenceDp(text1, text2, i+1, j, memo),
			longestCommonSubsequenceDp(text1, text2, i, j+1, memo),
		)
	}

	memo[i][j] = ans

	return ans
}

func LongestCommonSubsequenceV2(text1, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}
