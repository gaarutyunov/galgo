package sorting

import (
	"github.com/gaarutyunov/galgo/functional"
	"math/rand"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	r := rand.New(rand.NewSource(0))

	cases := []TestCase[int]{
		{
			name:     "all positive even",
			generate: generateInts(10, 10),
			expected: sort.Ints,
			sort:     SelectionSort[int],
		},
		{
			name:     "all positive odd",
			generate: generateInts(10, 11),
			expected: sort.Ints,
			sort:     SelectionSort[int],
		},
		{
			name:     "all negative even",
			generate: functional.Pipe2(generateInts(10, 10), toNegative),
			expected: sort.Ints,
			sort:     SelectionSort[int],
		},
		{
			name:     "all negative odd",
			generate: functional.Pipe2(generateInts(10, 11), toNegative),
			expected: sort.Ints,
			sort:     SelectionSort[int],
		},
		{
			name:     "random sign even",
			generate: functional.Pipe2(generateInts(10, 10), randomNegative(r, 30)),
			expected: sort.Ints,
			sort:     SelectionSort[int],
		},
		{
			name:     "random sign odd",
			generate: functional.Pipe2(generateInts(10, 11), randomNegative(r, 30)),
			expected: sort.Ints,
			sort:     SelectionSort[int],
		},
	}

	runCases(cases, t)
}
