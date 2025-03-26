package lm2033 // https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/

import (
	"slices"
)

func minOperations(grid [][]int, x int) int {
	size := len(grid) * len(grid[0])
	if size == 1 {
		return 0
	}

	counter := 0
	line := make([]int, 0, size)

	for m := 0; m < len(grid); m++ {
		line = append(line, grid[m]...)
	}

	slices.Sort(line)

	// middle element will be uni-value
	uni := line[size/2]

	for _, v := range line {
		delta := uni - v
		if delta < 0 {
			delta *= -1
		}

		if delta%x != 0 {
			counter = -1
			break
		}

		counter += delta / x
	}

	return counter
}

func Wrapper(grid [][]int, x int) int {
	return minOperations(grid, x)
}
