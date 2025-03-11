package searching

import "testing"

func TestBinarySearch(t *testing.T) {
	cases := []TestCase[int]{
		{
			input:    []int{1, 3, 6, 3, 7, 32, 67},
			expected: 3,
			search:   BinarySearch[int],
			target:   3,
		},
		{
			input:    []int{1, 3, 6, 7, 32, 67},
			expected: 1,
			search:   BinarySearch[int],
			target:   3,
		},
	}

	runCases("Binary Search", cases, t)
}
