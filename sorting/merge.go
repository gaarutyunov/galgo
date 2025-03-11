package sorting

import "cmp"

func MergeSort[T cmp.Ordered](arr []T) {
	helper := make([]T, len(arr), len(arr))
	mergeSort(arr, helper, 0, len(arr)-1)
}

func mergeSort[T cmp.Ordered](arr, helper []T, low, high int) {
	if low < high {
		middle := low + (high-low)/2
		mergeSort[T](arr, helper, low, middle)
		mergeSort[T](arr, helper, middle+1, high)
		merge(arr, helper, low, middle, high)
	}
}

func merge[T cmp.Ordered](arr, helper []T, low, middle, high int) {
	for i := low; i <= high; i++ {
		helper[i] = arr[i]
	}

	helperLeft := low
	helperRight := middle + 1
	current := low

	for helperLeft <= middle && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			arr[current] = helper[helperLeft]
			helperLeft++
		} else {
			arr[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	remaining := middle - helperLeft
	for i := 0; i <= remaining; i++ {
		arr[current+i] = helper[helperLeft+i]
	}
}
