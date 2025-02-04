// https://leetcode.com/problems/maximum-ascending-subarray-sum

package l_e_1800_array

func maxAscendingSum(nums []int) int {
	sum, sumMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			sum += nums[i]
		} else {
			if sum > sumMax {
				sumMax = sum
			}
			sum = nums[i]
		}
	}

	if sum > sumMax {
		return sum
	}

	return sumMax
}

func Wrapper(nums []int) int {
	return maxAscendingSum(nums)
}
