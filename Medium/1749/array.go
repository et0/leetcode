package lm1749 // https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/

func maxAbsoluteSum(nums []int) int {
	maxPlus, maxMinus := 0, 0
	max, min := 0, 0

	for _, v := range nums {
		// Kadane - find max with non-negative
		maxPlus += v
		if maxPlus < 0 {
			maxPlus = 0
		} else if maxPlus > max {
			max = maxPlus
		}

		// Kadane - find min with negative
		maxMinus += v
		if maxMinus > 0 {
			maxMinus = 0
		} else if maxMinus < min {
			min = maxMinus
		}
	}

	if min < 0 {
		min *= -1
	}

	if min > max {
		return min
	}
	return max
}

func Wrapper(nums []int) int {
	return maxAbsoluteSum(nums)
}
