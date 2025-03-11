package tree

import "cmp"

type Node[T cmp.Ordered] struct {
	value    T
	children []*Node[T]
}

func New[T cmp.Ordered](value T) *Node[T] {
	return &Node[T]{value: value}
}

func (n *Node[T]) Len() int {
	return len(n.children)
}

func (n *Node[T]) Value() T {
	return n.value
}

func (n *Node[T]) Children() []*Node[T] {
	return n.children
}

func (n *Node[T]) Child(i int) *Node[T] {
	if len(n.children) <= i {
		return nil
	}

	return n.children[i]
}

func (n *Node[T]) Append(value T) {
	n.children = append(n.children, New(value))
}

func (n *Node[T]) AppendNode(node *Node[T]) {
	n.children = append(n.children, node)
}

func (n *Node[T]) Remove(node *Node[T]) *Node[T] {
	for i, child := range n.children {
		if child == node {
			n.children = append(n.children[:i], n.children[i+1:]...)
			return node
		}
	}

	return nil
}
