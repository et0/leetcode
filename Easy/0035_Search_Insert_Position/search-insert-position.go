package searchinsertposition

import "fmt"

func searchInsert(nums []int, target int) int {
	size := len(nums)

	for i := 0; i < size; i++ {
		fmt.Println(i)
		if nums[i] > target {
			if i == 0 {
				return 0
			} else {
				return i
			}
		}
		if nums[i] == target {
			return i
		}
		if nums[i] < target && i+1 < size && nums[i+1] >= target {
			return i + 1
		}
		i++
	}
	return size
}

func searchInsert1(nums []int, target int) int {
	index := 0
	for _, v := range nums {
		if v >= target {
			break
		}
		index++
	}

	return index
}

// searchinsertposition.Wrapper([]int{1, 3, 5}, 5) // 0035
func Wrapper(nums []int, target int) int {
	return searchInsert(nums, target)
}
