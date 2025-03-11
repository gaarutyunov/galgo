package exercises

import (
	"github.com/gaarutyunov/galgo/binarytree"
	"github.com/gaarutyunov/galgo/stack"
)

func MaxDepthBinaryTreeV1(tree *binarytree.Node[int]) (val int) {
	var dfs func(tree *binarytree.Node[int], depth int)
	dfs = func(tree *binarytree.Node[int], depth int) {
		if tree == nil {
			return
		}
		val = max(val, depth)
		dfs(tree.Left, depth+1)
		dfs(tree.Right, depth+1)
	}

	dfs(tree, 1)

	return
}

type layer struct {
	node  *binarytree.Node[int]
	depth int
}

func MaxDepthBinaryTreeV2(tree *binarytree.Node[int]) (val int) {
	if tree == nil {
		return
	}
	s := stack.New[layer]()

	s.Push(layer{tree, 1})

	for !s.Empty() {
		l, ok := s.Pop()
		if !ok {
			return
		}

		val = max(val, l.depth)
		if l.node.Left != nil {
			s.Push(layer{l.node.Left, l.depth + 1})
		}
		if l.node.Right != nil {
			s.Push(layer{l.node.Right, l.depth + 1})
		}
	}

	return
}
