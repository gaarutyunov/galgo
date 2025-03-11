package sorting

import (
	"cmp"
	"github.com/gaarutyunov/galgo/searching"
)

func BinaryInsertionSort[T cmp.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		key := arr[i]

		loc := searching.BinarySearchRange(arr, key, 0, j, true)

		for ; j >= loc; j-- {
			arr[j+1] = arr[j]
		}

		arr[j+1] = key
	}
}
