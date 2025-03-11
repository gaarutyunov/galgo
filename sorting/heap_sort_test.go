package sorting

import "testing"

func TestHeapSort(t *testing.T) {
	cases := defaultIntCases(HeapSort)

	cases.Run(t)
}
