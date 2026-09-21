package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	var current uint64 = 0
	for current*current <= uint64(x) {
		current++
	}
	current--

	return int(current)
}

func main() {
	fmt.Println(mySqrt(8))
}
