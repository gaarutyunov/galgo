package sorting

import "github.com/gaarutyunov/galgo/numbers"

func RadixSort[T numbers.Integer](arr []T) {
	if len(arr) < 2 {
		return
	}

	m := maximum(arr)
	for exp := T(1); m/exp > 0; exp *= 10 {
		countSort(arr, exp, 0, len(arr)-1)
	}
}

func countSort[T numbers.Integer](arr []T, exp T, from, to int) {
	n := to - from + 1

	output := make([]T, n, n)
	var count [10]int

	for i := from; i <= to; i++ {
		count[(arr[i]/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := to; i >= from; i-- {
		output[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
	}

	copy(arr, output)
}

func maximum[T numbers.Integer](arr []T) (res T) {
	for _, v := range arr {
		if v > res {
			res = v
		}
	}

	return
}
