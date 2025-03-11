package binarytree

import "github.com/gaarutyunov/galgo/heap"

type Node[T comparable] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func New[T comparable](value T) *Node[T] {
	return &Node[T]{Value: value}
}

func FromHeapSlice[T comparable](arr []T) (root *Node[T]) {
	if len(arr) == 0 {
		return
	}

	root = New(arr[0])

	if len(arr) == 1 {
		return
	}

	nodes := make([]*Node[T], len(arr))
	nodes[0] = root

	var zero T

	for i := 0; i <= len(arr)/2-1; i++ {
		currNode := nodes[i]
		if currNode == nil {
			currNode = New(arr[i])
		}
		leftChild := heap.LeftChild(i)
		if leftChild < len(arr) && arr[leftChild] != zero {
			currNode.Left = New(arr[leftChild])
			nodes[leftChild] = currNode.Left
		}
		rightChild := heap.RightChild(i)
		if rightChild < len(arr) && arr[rightChild] != zero {
			currNode.Right = New(arr[rightChild])
			nodes[rightChild] = currNode.Right
		}
	}

	return
}
