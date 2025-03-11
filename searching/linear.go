package searching

import "cmp"

func LinearSearch[T cmp.Ordered](arr []T, target T) int {
	for i, t := range arr {
		if target == t {
			return i
		}
	}

	return -1
}
