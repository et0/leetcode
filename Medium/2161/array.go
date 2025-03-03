package lm2161 // https://leetcode.com/problems/partition-array-according-to-given-pivot/description/

func pivotArray(nums []int, pivot int) []int {
	size := len(nums)
	low, middle, high := make([]int, 0, size), make([]int, 0, size), make([]int, 0, size)

	for _, v := range nums {
		if v < pivot {
			low = append(low, v)
		} else if v > pivot {
			high = append(high, v)
		} else {
			middle = append(middle, v)
		}
	}

	nums = low
	nums = append(nums, middle...)
	nums = append(nums, high...)

	return nums

	// return slices.Concat(low, middle, high)
}

func Wrapper(nums []int, pivot int) []int {
	return pivotArray(nums, pivot)
}
