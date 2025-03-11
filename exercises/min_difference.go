package exercises

import (
	"github.com/gaarutyunov/galgo/numbers"
	"github.com/gaarutyunov/galgo/sorting"
)

// MinDifferenceV1 https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves
func MinDifferenceV1(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}

	sorting.QuickSort(nums)

	minDiff := numbers.MaxInt

	for left, right := 0, len(nums)-4; left < 4; left, right = left+1, right+1 {
		minDiff = min(minDiff, nums[right]-nums[left])
	}

	return minDiff
}
