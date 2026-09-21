package main // https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/

import (
	"slices"
)

func minimumDifference(nums []int, k int) int {
	minimum := 100000

	slices.Sort(nums)

	for left, right := 0, 0; left < len(nums); left++ {
		for ; right < len(nums) && right-left < k; right++ {

		}

		if right-left != k {
			continue
		}

		minimum = min(minimum, nums[right-1]-nums[left])
	}

	return minimum
}
