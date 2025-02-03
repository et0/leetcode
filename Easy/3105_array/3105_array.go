// https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/description

package l_e_3105_array

func longestMonotonicSubarray(nums []int) int {
	inc, dec, max := 1, 1, 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			inc++
		} else {
			inc = 1
		}
		if nums[i] > nums[i+1] {
			dec++
		} else {
			dec = 1
		}

		if inc > max {
			max = inc
		}
		if dec > max {
			max = dec
		}
	}

	return max
}

func Wrapper(nums []int) int {
	return longestMonotonicSubarray(nums)
}
