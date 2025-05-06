package le1920 // https://leetcode.com/problems/build-array-from-permutation/

func buildArray(nums []int) []int {
	new := make([]int, len(nums))
	for k, v := range nums {
		new[k] = nums[v]
	}
	return new
}

func Wrapper(nums []int) []int {
	return buildArray(nums)
}
