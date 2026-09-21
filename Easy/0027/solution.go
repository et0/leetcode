package main

func removeElement(nums []int, val int) int {
	left := 0
	for right := len(nums) - 1; left < right; left++ {
		if nums[left] != val {
			continue
		}

		for ; left < right && nums[right] == val; right-- {

		}
		nums[left], nums[right] = nums[right], nums[left]
	}

	return left
}
