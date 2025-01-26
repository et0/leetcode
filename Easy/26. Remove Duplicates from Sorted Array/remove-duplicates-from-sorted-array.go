package main

import "fmt"

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	unique := 0
	for i := 1; i < size; i++ {
		if nums[unique] == nums[i] {
			continue
		}
		unique++
		nums[unique] = nums[i]
	}

	return unique + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
