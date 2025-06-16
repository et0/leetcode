package le2016 // https://leetcode.com/problems/maximum-difference-between-increasing-elements/

func maximumDifference(nums []int) int {
	n := len(nums)
	copyN := make([]int, n)
	copy(copyN, nums)

	max := n - 1
	min := -1
	for i := max - 1; i >= 0; i-- {
		if copyN[i] < copyN[max] {
			copyN[i] = copyN[max]
		} else if copyN[i] > copyN[max] {
			max = i
		}

		try := copyN[max] - nums[i]
		if try != 0 && try > min {
			min = try
		}
	}

	return min
}

func Wrapper(nums []int) int {
	return maximumDifference(nums)
}
