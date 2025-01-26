package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	slice := make([]int, 0, 10)
	for del := 10; del <= 10000000000; del *= 10 {
		ost := x % del
		slice = append(slice, ost/(del/10))
		if ost >= x {
			break
		}
	}

	size := len(slice)
	for i := 0; i < size/2; i++ {
		if slice[i] != slice[size-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(123421))
}
