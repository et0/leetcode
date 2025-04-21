package le0283 // https://leetcode.com/problems/move-zeroes/description/

import "fmt"

func moveZeroes(nums []int) {
	fmt.Println(nums)
	one, two := 0, 0
	for ; two < len(nums); two++ {
		if nums[two] != 0 {
			nums[two], nums[one] = nums[one], nums[two]
			one++
		}
		fmt.Println(nums)
	}
}

func Wrapper(nums []int) {
	moveZeroes(nums)
}
