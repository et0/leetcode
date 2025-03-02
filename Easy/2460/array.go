package le2460 // https://leetcode.com/problems/apply-operations-to-an-array/description/

func applyOperations(nums []int) []int {
	size := len(nums)

	for i := 0; i < size-1; i++ {
		if nums[i] == 0 {
			copy := nums[:i]
			if i+1 < size {
				copy = append(copy, nums[i+1:]...)
			}
			copy = append(copy, 0)
			nums = copy

			i--
			size--

			continue
		}

		if nums[i] == nums[i+1] {
			nums[i] += nums[i+1]
			nums[i+1] = 0
		}
	}

	return nums
}

func Wrapper(nums []int) []int {
	return applyOperations(nums)
}
