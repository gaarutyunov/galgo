package sorting

import "cmp"

func QuickSort[T cmp.Ordered](arr []T) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort[T cmp.Ordered](arr []T, left, right int) {
	index := partition(arr, left, right)
	if left < index-1 {
		quickSort[T](arr, left, index-1)
	}
	if index < right {
		quickSort[T](arr, index, right)
	}
}

func partition[T cmp.Ordered](arr []T, left, right int) int {
	pivot := arr[left+(right-left)/2]

	for left <= right {
		for arr[left] < pivot {
			left++
		}

		for arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	return left
}
