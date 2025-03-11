package sorting

import (
	"cmp"
	"github.com/gaarutyunov/galgo/functional"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

type (
	TestCase[T cmp.Ordered] struct {
		name     string
		generate func(r *rand.Rand) [][]T
		sort     func([]T)
		expected func([]T)
	}

	TestCases[T cmp.Ordered] []TestCase[T]
)

func (c TestCases[T]) Run(t *testing.T) {
	runCases(c, t)
}

func runCases[T cmp.Ordered](cases []TestCase[T], t *testing.T) {
	r := rand.New(rand.NewSource(0))

	for _, caze := range cases {
		t.Run(caze.name, func(t *testing.T) {
			arr := caze.generate(r)
			expected := make([][]T, len(arr), len(arr))

			for i, a := range arr {
				expected[i] = make([]T, len(a), len(a))
				copy(expected[i], a)
				caze.expected(expected[i])
				caze.sort(a)
				assert.Equalf(t, expected[i], a, "case %d should equal", i)
			}
		})
	}
}

func generateInts(n, size int) func(r *rand.Rand) [][]int {
	return func(r *rand.Rand) (res [][]int) {
		res = make([][]int, n, n)

		for i := 0; i < n; i++ {
			res[i] = make([]int, size, size)

			for j := 0; j < size; j++ {
				res[i][j] = r.Intn(100)
			}
		}

		return
	}
}

func toNegative(arr [][]int) [][]int {
	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = -arr[i][j]
		}
	}

	return arr
}

func randomNegative(r *rand.Rand, p float32) func([][]int) [][]int {
	return func(arr [][]int) [][]int {
		for i := range arr {
			for j := range arr[i] {
				if int(p*100) < r.Intn(100) {
					arr[i][j] = -arr[i][j]
				}
			}
		}

		return arr
	}
}

func defaultIntCases(solution func([]int)) TestCases[int] {
	r := rand.New(rand.NewSource(1))

	return TestCases[int]{
		{
			name:     "all positive even",
			generate: generateInts(10, 10),
			expected: sort.Ints,
			sort:     solution,
		},
		{
			name:     "all positive odd",
			generate: generateInts(10, 11),
			expected: sort.Ints,
			sort:     solution,
		},
		{
			name:     "all negative even",
			generate: functional.Pipe2(generateInts(10, 10), toNegative),
			expected: sort.Ints,
			sort:     solution,
		},
		{
			name:     "all negative odd",
			generate: functional.Pipe2(generateInts(10, 11), toNegative),
			expected: sort.Ints,
			sort:     solution,
		},
		{
			name:     "random sign even",
			generate: functional.Pipe2(generateInts(10, 10), randomNegative(r, 30)),
			expected: sort.Ints,
			sort:     solution,
		},
		{
			name:     "random sign odd",
			generate: functional.Pipe2(generateInts(10, 11), randomNegative(r, 30)),
			expected: sort.Ints,
			sort:     solution,
		},
	}
}
