package exercises

import "github.com/gaarutyunov/galgo/heap"

// LastStoneWeightV1 https://leetcode.com/problems/last-stone-weight/
func LastStoneWeightV1(stones []int) int {
	h := heap.NewMax(stones)

	for !h.Empty() {
		y, ok := h.Pop()
		x, ok := h.Pop()
		if !ok {
			return y
		}

		if y != x {
			h.Push(y - x)
		}
	}

	return 0
}
