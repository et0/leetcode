package le3487 // https://leetcode.com/problems/maximum-unique-subarray-sum-after-deletion/

func maxSum(nums []int) int {
	uniq := make(map[int]bool)

	out := 0
	max := nums[0]
	if max >= 0 {
		uniq[max] = true
		out += max
	}

	size := len(nums)
	for i := 1; i < size; i++ {
		if nums[i] > max {
			max = nums[i]
		}

		if nums[i] < 0 {
			continue
		}

		if _, ok := uniq[nums[i]]; !ok {
			uniq[nums[i]] = true
			out += nums[i]
		}

	}

	if len(uniq) == 0 {
		return max
	}

	return out
}

func Wrapper(nums []int) int {
	return maxSum(nums)
}
