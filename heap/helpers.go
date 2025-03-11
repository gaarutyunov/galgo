package heap

func Parent(i int) int {
	return (i - 1) >> 1
}

func LeftChild(i int) int {
	return 2*i + 1
}

func RightChild(i int) int {
	return 2*i + 2
}
