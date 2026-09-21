package main

import "fmt"

func plusOne(digits []int) []int {
	size := len(digits)
	for i := size - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			break
		}
		if i == 0 {
			digits[i] = 1
			digits = append(digits, 0)
		} else {
			digits[i] = 0
		}
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{9, 2, 9}))
}
