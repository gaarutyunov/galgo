package exercises

import (
	"github.com/gaarutyunov/galgo/exercises"
	"testing"
)

func TestLastStoneWeightV1(t *testing.T) {
	cases := TestCases[[]int, int]{
		{
			Input:    []int{2, 7, 4, 1, 8, 1},
			Expected: 1,
			Solution: exercises.LastStoneWeightV1,
		},
		{
			Input:    []int{1},
			Expected: 1,
			Solution: exercises.LastStoneWeightV1,
		},
		{
			Input:    []int{9, 10, 4, 5, 7, 1},
			Expected: 0,
			Solution: exercises.LastStoneWeightV1,
		},
		{
			Input:    []int{10, 5, 4, 10, 3, 1, 7, 8},
			Expected: 0,
			Solution: exercises.LastStoneWeightV1,
		},
	}

	cases.Run("TestLastStoneWeightV1", t)
}
