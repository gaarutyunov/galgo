package exercises

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type (
	testCase[T any, R any] struct {
		Input     T
		Solution  func(T) R
		SolutionI func(int) func(T) R
		Expected  R
	}

	TestCases[T any, R any] []testCase[T, R]
)

func runCases[T any, R any](name string, cases []testCase[T, R], t *testing.T) {
	for i, caze := range cases {
		t.Run(name+" "+strconv.Itoa(i), func(t *testing.T) {
			var solution func(T) R
			if caze.SolutionI != nil {
				solution = caze.SolutionI(i)
			} else {
				solution = caze.Solution
			}
			actual := solution(caze.Input)
			assert.Equal(t, caze.Expected, actual)
		})
	}
}

func (r TestCases[T, R]) Run(name string, t *testing.T) {
	runCases(name, r, t)
}
