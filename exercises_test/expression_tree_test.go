package exercises

import (
	"github.com/gaarutyunov/galgo/exercises"
	"testing"
)

func TestExpressionTreeV1(t *testing.T) {
	solution := exercises.ExpressionTreeSolution(exercises.BuilderFunc(exercises.ExpressionTreeV1))

	cases := TestCases[[]string, int]{
		{
			Input:    []string{"3", "4", "+", "2", "*", "7", "/"},
			Solution: solution,
			Expected: 2,
		},
		{
			Input:    []string{"4", "5", "2", "7", "+", "-", "*"},
			Solution: solution,
			Expected: -16,
		},
	}

	cases.Run("TestExpressionTreeV1", t)
}
