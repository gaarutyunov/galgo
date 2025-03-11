package exercises

import (
	"cmp"
	"github.com/gaarutyunov/galgo/linkedlist"
)

func RemoveDuplicatesV1[T cmp.Ordered](list *linkedlist.Node[T]) *linkedlist.Node[T] {
	p1 := list
	p2 := p1.Next()
	prev := p1

	for p1.HasNext() {
		if p1.Value() == p2.Value() {
			prev.SkipNext()
			p2 = prev.Next()
		} else if !p2.HasNext() {
			p1 = p1.Next()
			p2 = p1.Next()
			prev = p1
		} else {
			prev = p2
			p2 = p2.Next()
		}
	}

	return list
}

func KthToLast[T cmp.Ordered](k int) func(list *linkedlist.Node[T]) *linkedlist.Node[T] {
	return func(list *linkedlist.Node[T]) *linkedlist.Node[T] {
		var counter int

		for ; list.HasNext() && counter != k; counter++ {
			list = list.Next()
		}

		return list
	}
}

func DeleteMiddle[T cmp.Ordered](middle *linkedlist.Node[T]) func(list *linkedlist.Node[T]) *linkedlist.Node[T] {
	return func(list *linkedlist.Node[T]) *linkedlist.Node[T] {
		p1 := list

		for p1.Next() != middle {
			p1 = p1.Next()
		}

		p1.SkipNext()

		return list
	}
}

func Partition[T cmp.Ordered](x T) func(list *linkedlist.Node[T]) *linkedlist.Node[T] {
	return func(list *linkedlist.Node[T]) *linkedlist.Node[T] {
		var lHead, lTail, hHead, hTail *linkedlist.Node[T]

		for list != nil {
			if list.Value() < x {
				var lNode *linkedlist.Node[T]
				lNode, list = list.Detach()
				if lTail == nil {
					lTail = lNode
					lHead = lTail
				} else {
					lTail = lTail.AppendNode(lNode)
				}
			} else {
				var hNode *linkedlist.Node[T]
				hNode, list = list.Detach()
				if hTail == nil {
					hTail = hNode
					hHead = hTail
				} else {
					hTail = hTail.AppendNode(hNode)
				}
			}
		}

		lTail.AppendNode(hHead)

		return lHead
	}
}
