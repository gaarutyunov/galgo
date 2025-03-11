package exercises

// MinCostClimbingStairsV1 https://leetcode.com/problems/min-cost-climbing-stairs
func MinCostClimbingStairsV1(cost []int) int {
	memo := make([]int, len(cost)+1)
	for i := range memo {
		memo[i] = -1
	}

	return minCostClimbingStairs(cost, len(cost), memo)
}

func minCostClimbingStairs(costs []int, i int, memo []int) int {
	if v := memo[i]; v != -1 {
		return v
	}

	memo[i] = min(minCostClimbingStairs(costs, i-1, memo)+costs[i-1], minCostClimbingStairs(costs, i-2, memo)+costs[i-2])

	return memo[i]
}

func MinCostClimbingStairsV2(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}

func MinCostClimbingStairsV3(cost []int) int {
	prevPrev, prev := 0, 0

	for i := 2; i <= len(cost); i++ {
		newVal := min(prev+cost[i-1], prevPrev+cost[i-2])
		prevPrev = prev
		prev = newVal
	}

	return prev
}
