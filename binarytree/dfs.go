package binarytree

import (
	"github.com/gaarutyunov/galgo/stack"
	"iter"
)

func InorderDFS[T comparable](tree *Node[T]) iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		s := stack.New[*Node[T]]()
		root := tree

		for !s.Empty() || root != nil {
			if root != nil {
				s.Push(root)
				root = root.Left
			} else {
				root, _ = s.Pop()
				if !yield(root) {
					return
				}
				root = root.Right
			}
		}
	}
}
