package exercises

import (
	"github.com/gaarutyunov/galgo/binarytree"
	"github.com/gaarutyunov/galgo/queue"
)

type zigzagLayer struct {
	node  *binarytree.Node[int]
	depth int
	ltr   bool
}

func ZigzagTraversalV1(tree *binarytree.Node[int]) (res [][]int) {
	q := queue.New[zigzagLayer]()

	q.Push(zigzagLayer{tree, 1, true})
	currDepth := 1
	var currRow []int

	for !q.Empty() {
		l, ok := q.Pop()
		if !ok {
			return
		}
		if currDepth == l.depth {
			currRow = append(currRow, l.node.Value)
		} else {
			res = append(res, currRow)
			currDepth = l.depth
			currRow = []int{l.node.Value}
		}
		if l.ltr {
			if l.node.Left != nil {
				q.Push(zigzagLayer{l.node.Left, l.depth + 1, !l.ltr})
			}
			if l.node.Right != nil {
				q.Push(zigzagLayer{l.node.Right, l.depth + 1, !l.ltr})
			}
		} else {
			if l.node.Right != nil {
				q.Push(zigzagLayer{l.node.Right, l.depth + 1, !l.ltr})
			}
			if l.node.Left != nil {
				q.Push(zigzagLayer{l.node.Left, l.depth + 1, !l.ltr})
			}
		}
	}

	return
}
