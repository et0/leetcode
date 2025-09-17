package main

import "fmt"

func removeElement(nums []int, val int) int {
	fmt.Println(nums)
	left := 0
	for right := len(nums) - 1; left < right; left++ {
		if nums[left] != val {
			continue
		}

		for ; left < right && nums[right] == val; right-- {

		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	fmt.Println(nums)
	return left
}

func main() {
	fmt.Println(removeElement([]int{1, 2, 3, 4, 2, 0, 3}, 2))
}
