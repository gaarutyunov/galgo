package heap

import (
	"cmp"
)

type Heap[T cmp.Ordered] interface {
	Empty() bool
	Len() int
	Cap() int
	Push(v T)
	Pop() (T, bool)
	Peek() (T, bool)
	Set(i int, v T)
}

type baseHeap[T cmp.Ordered] struct {
	arr []T
}

func (h *baseHeap[T]) Len() int {
	return len(h.arr)
}

func (h *baseHeap[T]) Empty() bool {
	return len(h.arr) == 0
}

func (h *baseHeap[T]) Cap() int {
	return cap(h.arr)
}
