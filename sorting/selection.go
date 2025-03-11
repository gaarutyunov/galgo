package sorting

import "cmp"

func SelectionSort[T cmp.Ordered](arr []T) {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if i == minIdx {
			continue
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
