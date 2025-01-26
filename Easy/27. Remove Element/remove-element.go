package main

import "fmt"

func removeElement1(nums []int, val int) int {
	fmt.Println(nums)
	writeIndex := 0
	for readIndex := 0; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != val {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}
	fmt.Println(nums)
	return writeIndex
}

func removeElement(nums []int, val int) int {
	fmt.Println(nums)

	size := len(nums)
	if size == 0 {
		return 0
	}

	unique := 0
	for i := 0; i < size; i++ {
		if nums[i] != val {
			unique++
			continue
		}

		for j := size - 1; j > i; j-- {
			size--
			if nums[j] == val {
				continue
			}
			nums[i], nums[j] = nums[j], nums[i]
			unique++
			break
		}
	}

	fmt.Println(nums)

	return unique
}

func main() {
	fmt.Println(removeElement1([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
