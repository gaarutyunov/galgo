package exercises

func ClimbStairsV1(n int) int {
	memo := make([]int, n+1)

	for i := range memo {
		memo[i] = -1
	}

	return climbStairs(n, memo)
}

func climbStairs(n int, memo []int) int {
	if n <= 0 {
		return 0
	}

	if memo[n] != -1 {
		return memo[n]
	}

	memo[n] = climbStairs(n-1, memo) + climbStairs(n-2, memo)

	return memo[n]
}

func ClimbStairsV2(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)

	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func ClimbStairsV3(n int) int {
	if n <= 2 {
		return n
	}

	prev, current := 1, 2
	for i := 3; i <= n; i++ {
		next := prev + current
		prev = current
		current = next
	}

	return current
}
