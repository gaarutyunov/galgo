package exercises

// HouseRobberV1 https://leetcode.com/problems/house-robber
func HouseRobberV1(nums []int) int {
	memo := make([]int, len(nums))

	for i := range memo {
		memo[i] = -1
	}

	return houseRobber(nums, len(nums)-1, memo)
}

func houseRobber(nums []int, i int, memo []int) int {
	if i == 0 {
		return nums[0]
	}

	if i == 1 {
		return max(nums[0], nums[1])
	}

	if memo[i] != -1 {
		return memo[i]
	}

	memo[i] = max(houseRobber(nums, i-1, memo), houseRobber(nums, i-2, memo)+nums[i])

	return memo[i]
}

func HouseRobberV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

func HouseRobberV3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	prevPrev, prev := nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		newVal := max(prev, prevPrev+nums[i])
		prevPrev = prev
		prev = newVal
	}

	return prev
}
