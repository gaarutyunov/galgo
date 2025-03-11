package linkedlist

type Node[T any] struct {
	value T
	next  *Node[T]
}

func New[T any](value T) *Node[T] {
	return &Node[T]{value: value}
}

func FromSlice[T any](values []T) (list *Node[T]) {
	list = New(values[0])
	p := list

	for _, value := range values[1:] {
		p.next = New(value)
		p = p.Next()
	}

	return
}

func (n *Node[T]) ToSlice() (slice []T) {
	p := n

	for p != nil {
		slice = append(slice, p.Value())
		p = p.Next()
	}

	return
}

func ToSlice[T any](n *Node[T]) []T {
	return n.ToSlice()
}

func (n *Node[T]) Len() int {
	node := n
	res := 1

	for node.HasNext() {
		res += 1
		node = n.next
	}

	return res
}

func (n *Node[T]) Value() T {
	return n.value
}

func (n *Node[T]) Detach() (node *Node[T], rest *Node[T]) {
	rest = n.next
	n.next = nil
	node = n
	return
}

func (n *Node[T]) HasNext() bool {
	return n.next != nil
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Append(v T) *Node[T] {
	return n.AppendNode(New(v))
}

func (n *Node[T]) AppendNode(node *Node[T]) *Node[T] {
	if !n.HasNext() {
		n.next = node
	} else {
		n.next.AppendNode(node)
	}

	return node
}

func (n *Node[T]) SkipNext() (*Node[T], bool) {
	if !n.HasNext() {
		return nil, false
	}

	oldNext := n.next

	n.next = n.next.Next()

	return oldNext, true
}

func (n *Node[T]) ReplaceNextNode(node *Node[T]) {
	if !n.HasNext() {
		n.AppendNode(node)
		return
	}

	node.next = n.next.Next()
	n.next = node
}

func (n *Node[T]) ReplaceNext(v T) {
	node := New(v)
	n.ReplaceNextNode(node)
}

func (n *Node[T]) Insert(node *Node[T]) {
	oldNext := n.next
	n.next = node
	node.next = oldNext
}
