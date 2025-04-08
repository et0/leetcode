package le3396 // https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct/description/

func minimumOperations(nums []int) int {
	all := make(map[int]bool)

	size := len(nums)
	for i := size - 1; i >= 0; i-- {
		if !all[nums[i]] {
			all[nums[i]] = true
			continue
		}

		i++
		if i%3 == 0 {
			return i / 3
		}
		return i/3 + 1
	}

	return 0
}

func Wrapper(nums []int) int {
	return minimumOperations(nums)
}
