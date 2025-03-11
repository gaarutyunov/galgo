package sorting

import (
	"cmp"
	"github.com/gaarutyunov/galgo/heap"
)

func HeapSort[T cmp.Ordered](arr []T) {
	heap.Max(arr)

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heap.MaxDown(arr, 0, i)
	}
}
