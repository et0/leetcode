package main // https://leetcode.com/problems/minimum-size-subarray-sum/

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	minLength := 100001
	sum := 0

	for left, right := 0, 0; left < len(nums); left++ {
		for ; right < len(nums) && sum < target; right++ {
			sum += nums[right]
		}

		if sum >= target {
			minLength = min(minLength, right-left)
		}

		sum -= nums[left]
	}

	if minLength == 100001 {
		return 0
	}

	return minLength
}

func main() {
	fmt.Println(minSubArrayLen(5, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
