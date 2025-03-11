package exercises

import "github.com/gaarutyunov/galgo/searching"

// LongestIncreasingSubsequenceV1 https://leetcode.com/problems/longest-increasing-subsequence
func LongestIncreasingSubsequenceV1(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	var res int

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				res = max(res, dp[i])
			}
		}
	}

	return res
}

func LongestIncreasingSubsequenceV2(nums []int) int {
	var sub []int

	for _, num := range nums {
		if len(sub) == 0 || sub[len(sub)-1] < num {
			sub = append(sub, num)
		} else {
			idx := searching.BinarySearchRange(sub, num, 0, len(sub)-1, true)
			sub[idx] = num
		}
	}

	return len(sub)
}
