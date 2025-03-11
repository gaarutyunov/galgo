package sorting

import "cmp"

func BubbleSort[T cmp.Ordered](arr []T) {
	for i := 0; i < len(arr); i++ {
		var swapped bool
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
