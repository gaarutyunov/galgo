package searching

import "cmp"

func BinarySearchRange[T cmp.Ordered](arr []T, target T, low, high int, returnLow bool) int {
	for low <= high {
		mid := low + (high-low)/2
		switch {
		case target == arr[mid]:
			return mid
		case target > arr[mid]:
			low = mid + 1
		default:
			high = mid - 1
		}
	}

	if !returnLow {
		return -1
	}

	return low
}

func BinarySearch[T cmp.Ordered](arr []T, target T) int {
	return BinarySearchRange(arr, target, 0, len(arr)-1, false)
}
