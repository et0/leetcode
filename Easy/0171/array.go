package le0171 // https://leetcode.com/problems/excel-sheet-column-number/description/

import (
	"math"
)

func titleToNumber(columnTitle string) int {
	length := len(columnTitle) - 1
	number := 0
	for _, v := range columnTitle {
		if length == 0 {
			number += int(v - 64)
		} else {
			number += (int(v) - 64) * int(math.Pow(26.0, float64(length)))
		}
		length--
	}

	return number
}

func Wrapper(columnTitle string) int {
	return titleToNumber(columnTitle)
}
