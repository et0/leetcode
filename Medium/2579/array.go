package lm2579 // https://leetcode.com/problems/count-total-number-of-colored-cells/description/

func coloredCells(n int) int64 {
	var result int64 = 0

	for i := 1; i < n; i++ {
		result += int64(i*2 - 1)
	}
	result *= 2
	result += int64(n*2 - 1)

	return result
}

func Wrapper(n int) int64 {
	return coloredCells(n)
}
