package main // https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		if numbers[left]+numbers[right] == target {
			break
		}

		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return []int{left + 1, right + 1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
