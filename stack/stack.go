package stack

type (
	Stack[T any] struct {
		head *node[T]
	}

	node[T any] struct {
		value T
		next  *node[T]
	}
)

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func newNode[T any](val T) *node[T] {
	return &node[T]{value: val}
}

func (n *node[T]) Value() T {
	return n.value
}

func (n *Stack[T]) Push(val T) {
	node := newNode(val)
	node.next = n.head
	n.head = node
}

func (n *Stack[T]) Pop() (T, bool) {
	if n.Empty() {
		var zero T
		return zero, false
	}

	head := n.head
	n.head = n.head.next
	return head.value, true
}

func (n *Stack[T]) Peek() (T, bool) {
	if n.Empty() {
		var zero T
		return zero, false
	}

	return n.head.value, true
}

func (n *Stack[T]) Empty() bool {
	return n.head == nil
}

func (n *Stack[T]) Head() (T, bool) {
	if n.Empty() {
		var zero T
		return zero, false
	}
	return n.head.value, true
}
