package exercises

import (
	"github.com/gaarutyunov/galgo/exercises"
	"testing"
)

type CoinChangeArgs struct {
	Coins  []int
	Amount int
}

func TestCoinChangeV2(t *testing.T) {
	solution := func(args CoinChangeArgs) int {
		return exercises.CoinChangeV2(args.Coins, args.Amount)
	}

	cases := []testCase[CoinChangeArgs, int]{
		{
			Input:    CoinChangeArgs{Coins: []int{1}, Amount: 1},
			Solution: solution,
			Expected: 1,
		},
		{
			Input:    CoinChangeArgs{Coins: []int{2}, Amount: 3},
			Solution: solution,
			Expected: -1,
		},
		{
			Input:    CoinChangeArgs{Coins: []int{2}, Amount: 0},
			Solution: solution,
			Expected: 0,
		},
		{
			Input:    CoinChangeArgs{Coins: []int{1, 2, 5}, Amount: 100},
			Solution: solution,
			Expected: 20,
		},
	}

	runCases("Coin Change V2", cases, t)
}
