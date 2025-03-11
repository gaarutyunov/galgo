package sorting

import (
	"math/rand"
	"sort"
	"testing"
)

func TestRadixSort(t *testing.T) {
	cases := []TestCase[int]{
		{
			name:     "all positive",
			generate: generateInts(10, 10),
			sort:     RadixSort[int],
			expected: sort.Ints,
		},
		{
			name: "empty",
			generate: func(r *rand.Rand) [][]int {
				return [][]int{
					{},
				}
			},
			sort:     RadixSort[int],
			expected: sort.Ints,
		},
		{
			name:     "one",
			generate: generateInts(1, 1),
			sort:     RadixSort[int],
			expected: sort.Ints,
		},
	}

	runCases(cases, t)
}
