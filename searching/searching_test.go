package searching

import (
	"cmp"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type TestCase[T cmp.Ordered] struct {
	input    []T
	target   T
	search   func([]T, T) int
	expected int
}

func runCases[T cmp.Ordered](name string, cases []TestCase[T], t *testing.T) {
	for i, caze := range cases {
		t.Run(name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := caze.search(caze.input, caze.target)
			assert.Equal(t, caze.expected, actual)
		})
	}
}
