package sorting

import "cmp"

func InsertionSort[T cmp.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}
