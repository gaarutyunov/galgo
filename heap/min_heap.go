package heap

import "cmp"

type minHeap[T cmp.Ordered] struct {
	*baseHeap[T]
}

func NewMin[T cmp.Ordered](arr []T) Heap[T] {
	h := &minHeap[T]{
		&baseHeap[T]{
			arr: arr,
		},
	}

	h.heapify()

	return h
}

func (h *minHeap[T]) Push(v T) {
	h.arr = append(h.arr, v)
	i := len(h.arr) - 1

	h.up(i)
}

func (h *minHeap[T]) Set(i int, v T) {
	if i > len(h.arr)-1 || h.arr[i] == v {
		return
	}
	oldVal := h.arr[i]

	h.arr[i] = v

	if oldVal < v {
		h.down(i)
	} else {
		h.up(i)
	}
}

func (h *minHeap[T]) Peek() (res T, ok bool) {
	if h.Empty() {
		return res, false
	}

	return h.arr[0], true
}

func (h *minHeap[T]) Pop() (res T, ok bool) {
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

func (h *minHeap[T]) heapify() {
	if len(h.arr) == 0 {
		return
	}
	Min(h.arr)
}

func (h *minHeap[T]) down(i int) {
	MinDown(h.arr, i, h.Len())
}

func (h *minHeap[T]) up(i int) {
	MinUp(h.arr, i, 0)
}

func Min[T cmp.Ordered](arr []T) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		MinDown(arr, i, len(arr))
	}
}

func MinDown[T cmp.Ordered](arr []T, root, end int) {
	left := LeftChild(root)
	right := RightChild(root)
	smallest := root
	if left < end && arr[left] < arr[smallest] {
		smallest = left
	}
	if right < end && arr[right] < arr[smallest] {
		smallest = right
	}
	if smallest != root {
		arr[root], arr[smallest] = arr[smallest], arr[root]
		MinDown[T](arr, smallest, end)
	}
}

func MinUp[T cmp.Ordered](arr []T, start, root int) {
	for start > root && arr[start] < arr[Parent(start)] {
		arr[start], arr[Parent(start)] = arr[Parent(start)], arr[start]
		start = Parent(start)
	}
}
