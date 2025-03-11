package heap

import "cmp"

type maxHeap[T cmp.Ordered] struct {
	*baseHeap[T]
}

func NewMax[T cmp.Ordered](arr []T) Heap[T] {
	h := &maxHeap[T]{
		&baseHeap[T]{
			arr: arr,
		},
	}

	h.heapify()

	return h
}

func (h *maxHeap[T]) Push(v T) {
	h.arr = append(h.arr, v)

	h.up(len(h.arr) - 1)
}

func (h *maxHeap[T]) Set(i int, v T) {
	if i > len(h.arr)-1 || h.arr[i] == v {
		return
	}
	oldVal := h.arr[i]

	h.arr[i] = v

	if oldVal > v {
		h.down(i)
	} else {
		h.up(i)
	}
}

func (h *maxHeap[T]) Peek() (res T, ok bool) {
	if h.Empty() {
		return res, false
	}

	return h.arr[0], true
}

func (h *maxHeap[T]) Pop() (res T, ok bool) {
	if h.Empty() {
		return res, false
	}

	if h.Len() == 1 {
		root := h.arr[0]
		h.arr = h.arr[:0]
		return root, true
	}

	root := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.down(0)

	return root, true
}

func (h *maxHeap[T]) heapify() {
	if len(h.arr) == 0 {
		return
	}
	Max(h.arr)
}

func (h *maxHeap[T]) down(i int) {
	MaxDown(h.arr, i, h.Len())
}

func (h *maxHeap[T]) up(i int) {
	MaxUp(h.arr, i, 0)
}

func Max[T cmp.Ordered](arr []T) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		MaxDown(arr, i, len(arr))
	}
}

func MaxDown[T cmp.Ordered](arr []T, root, end int) {
	left := LeftChild(root)
	right := RightChild(root)
	largest := root
	if left < end && arr[left] > arr[largest] {
		largest = left
	}
	if right < end && arr[right] > arr[largest] {
		largest = right
	}
	if largest != root {
		arr[root], arr[largest] = arr[largest], arr[root]
		MaxDown[T](arr, largest, end)
	}
}

func MaxUp[T cmp.Ordered](arr []T, start, root int) {
	for start > root && arr[start] > arr[Parent(start)] {
		arr[start], arr[Parent(start)] = arr[Parent(start)], arr[start]
		start = Parent(start)
	}
}
