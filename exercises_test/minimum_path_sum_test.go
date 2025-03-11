package exercises

import (
	"github.com/gaarutyunov/galgo/exercises"
	"testing"
)

func TestMinimumPathSumV1(t *testing.T) {
	cases := []testCase[[][]int, int]{
		{
			Input: [][]int{
				{1, 2, 5},
				{3, 2, 1},
			},
			Solution: exercises.MinimumPathSumV1,
			Expected: 6,
		},
	}

	runCases("Minimum Path Sum V1", cases, t)
}

func TestMinimumPathSumV2(t *testing.T) {
	cases := []testCase[[][]int, int]{
		{
			Input: [][]int{
				{1, 2, 5},
				{3, 2, 1},
			},
			Solution: exercises.MinimumPathSumV2,
			Expected: 6,
		},
		{
			Input: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			Solution: exercises.MinimumPathSumV2,
			Expected: 7,
		},
	}

	runCases("Minimum Path Sum V2", cases, t)
}
