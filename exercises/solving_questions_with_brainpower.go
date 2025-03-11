package exercises

// SolvingQuestionsWithBrainpowerV1 https://leetcode.com/problems/solving-questions-with-brainpower/
func SolvingQuestionsWithBrainpowerV1(questions [][]int) int {
	memo := make([]int, len(questions))

	return solvingQuestionsWithBrainpower(questions, 0, memo)
}

func solvingQuestionsWithBrainpower(questions [][]int, i int, memo []int) int {
	if i >= len(questions) {
		return 0
	}

	if memo[i] != -1 {
		return memo[i]
	}

	j := i + questions[i][1] + 1

	memo[i] = max(
		questions[i][0]+solvingQuestionsWithBrainpower(questions, j, memo),
		solvingQuestionsWithBrainpower(questions, i+1, memo),
	)

	return memo[i]
}

func SolvingQuestionsWithBrainpowerV2(questions [][]int) int {
	n := len(questions)

	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		j := i + questions[i][1] + 1
		dp[i] = max(questions[i][0]+dp[min(j, n)], dp[i+1])
	}

	return dp[0]
}
