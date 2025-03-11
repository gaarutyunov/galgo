package exercises

import (
	"github.com/gaarutyunov/galgo/exercises"
	"testing"
)

func TestUniquePathObstacleV1(t *testing.T) {
	cases := TestCases[[][]int, int]{
		{
			Input: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			Expected: 2,
			Solution: exercises.UniquePathObstacleV1,
		},
		{
			Input: [][]int{
				{0, 0},
				{0, 1},
			},
			Expected: 0,
			Solution: exercises.UniquePathObstacleV1,
		},
		{
			Input: [][]int{
				{0, 1},
				{0, 0},
			},
			Expected: 1,
			Solution: exercises.UniquePathObstacleV1,
		},
		{
			Input: [][]int{
				{1},
			},
			Expected: 0,
			Solution: exercises.UniquePathObstacleV1,
		},
		{
			Input: [][]int{
				{0, 0},
				{1, 0},
			},
			Expected: 1,
			Solution: exercises.UniquePathObstacleV1,
		},
		{
			Input: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			Expected: 7,
			Solution: exercises.UniquePathObstacleV1,
		},
	}

	cases.Run("UniquePathObstacleV1", t)
}

func TestUniquePathObstacleV2(t *testing.T) {
	cases := TestCases[[][]int, int]{
		{
			Input: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			Expected: 2,
			Solution: exercises.UniquePathObstacleV2,
		},
		{
			Input: [][]int{
				{0, 0},
				{0, 1},
			},
			Expected: 0,
			Solution: exercises.UniquePathObstacleV2,
		},
		{
			Input: [][]int{
				{0, 1},
				{0, 0},
			},
			Expected: 1,
			Solution: exercises.UniquePathObstacleV2,
		},
		{
			Input: [][]int{
				{1},
			},
			Expected: 0,
			Solution: exercises.UniquePathObstacleV2,
		},
		{
			Input: [][]int{
				{0, 0},
				{1, 0},
			},
			Expected: 1,
			Solution: exercises.UniquePathObstacleV2,
		},
		{
			Input: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			Expected: 7,
			Solution: exercises.UniquePathObstacleV2,
		},
	}

	cases.Run("UniquePathObstacleV2", t)
}

func TestUniquePathObstacleV3(t *testing.T) {
	cases := TestCases[[][]int, int]{
		{
			Input: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			Expected: 2,
			Solution: exercises.UniquePathObstacleV3,
		},
		{
			Input: [][]int{
				{0, 0},
				{0, 1},
			},
			Expected: 0,
			Solution: exercises.UniquePathObstacleV3,
		},
		{
			Input: [][]int{
				{0, 1},
				{0, 0},
			},
			Expected: 1,
			Solution: exercises.UniquePathObstacleV3,
		},
		{
			Input: [][]int{
				{1},
			},
			Expected: 0,
			Solution: exercises.UniquePathObstacleV3,
		},
		{
			Input: [][]int{
				{0, 0},
				{1, 0},
			},
			Expected: 1,
			Solution: exercises.UniquePathObstacleV3,
		},
		{
			Input: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			Expected: 7,
			Solution: exercises.UniquePathObstacleV3,
		},
	}

	cases.Run("UniquePathObstacleV3", t)
}
