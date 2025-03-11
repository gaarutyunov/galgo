package exercises

import (
	"github.com/gaarutyunov/galgo/binarytree"
)

func MinDifferenceTreeV1(tree *binarytree.Node[int]) int {
	var i, prev, ans int

	for node := range binarytree.InorderDFS(tree) {
		if i == 0 {
			prev = node.Value
		} else {
			ans = min(ans, node.Value-prev)
		}
		i++
	}

	return ans
}
