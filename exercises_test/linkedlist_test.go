package exercises

import (
	exercises2 "github.com/gaarutyunov/galgo/exercises"
	"github.com/gaarutyunov/galgo/functional"
	"github.com/gaarutyunov/galgo/linkedlist"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	solution := functional.Pipe5(
		functional.StringToRuneSlice,
		linkedlist.FromSlice,
		exercises2.RemoveDuplicatesV1,
		linkedlist.ToSlice,
		functional.RuneSliceToString,
	)

	cases := []testCase[string, string]{
		{
			Input:    "FOLLOW UP",
			Solution: solution,
			Expected: "FOLW UP",
		},
		{
			Input:    "T",
			Solution: solution,
			Expected: "T",
		},
		{
			Input:    "TTTT",
			Solution: solution,
			Expected: "T",
		},
	}

	runCases("Remove Duplicates", cases, t)
}

func TestKthToLast(t *testing.T) {
	solution := functional.Pipe5(
		functional.StringToRuneSlice,
		linkedlist.FromSlice,
		exercises2.KthToLast[rune](1),
		linkedlist.ToSlice,
		functional.RuneSliceToString,
	)

	cases := []testCase[string, string]{
		{
			Input:    "1234",
			Solution: solution,
			Expected: "234",
		},
		{
			Input:    "1",
			Solution: solution,
			Expected: "1",
		},
	}

	runCases("KthToLast", cases, t)
}

func TestDeleteMiddle(t *testing.T) {
	middles := []*linkedlist.Node[rune]{
		linkedlist.New[rune]('5'),
	}

	solution := func(i int) func(*linkedlist.Node[rune]) string {
		return functional.Pipe3(
			exercises2.DeleteMiddle(middles[i]),
			linkedlist.ToSlice,
			functional.RuneSliceToString,
		)
	}

	cases := []testCase[*linkedlist.Node[rune], string]{
		{
			Input: func() *linkedlist.Node[rune] {
				l := linkedlist.New[rune]('1')

				l.Append('2').
					AppendNode(middles[0]).
					Append('3')

				return l
			}(),
			SolutionI: solution,
			Expected:  "123",
		},
	}

	runCases("Delete Middle", cases, t)
}

func TestPartition(t *testing.T) {
	x := []int{
		5,
	}

	solution := func(i int) func([]int) []int {
		return functional.Pipe3(
			linkedlist.FromSlice,
			exercises2.Partition(x[i]),
			linkedlist.ToSlice,
		)
	}

	cases := []testCase[[]int, []int]{
		{
			Input:     []int{3, 5, 8, 5, 10, 2, 1},
			SolutionI: solution,
			Expected:  []int{3, 2, 1, 5, 8, 5, 10},
		},
	}

	runCases("Partition", cases, t)
}
