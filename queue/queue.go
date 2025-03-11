package queue

type (
	Queue[T any] struct {
		head, tail *node[T]
	}

	node[T any] struct {
		value T
		next  *node[T]
	}
)

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func newNode[T any](value T) *node[T] {
	return &node[T]{value: value}
}

func (q *Queue[T]) Push(value T) {
	node := newNode(value)
	if q.Empty() {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

func (q *Queue[T]) Pop() (T, bool) {
	if q.Empty() {
		var zero T
		return zero, false
	}
	head := q.head
	if head == q.tail {
		q.tail = nil
	}
	q.head = q.head.next
	return head.value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.Empty() {
		var zero T
		return zero, false
	}

	return q.head.value, true
}

func (q *Queue[T]) Empty() bool {
	return q.head == nil
}
