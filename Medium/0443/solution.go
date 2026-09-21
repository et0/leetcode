package main // https://leetcode.com/problems/string-compression/

import (
	"strconv"
)

func compress(chars []byte) int {
	length := len(chars)
	var idx int

	for left, right := 0, 0; left < length; left = right {

		for ; right < length && chars[left] == chars[right]; right++ {

		}

		chars[idx] = chars[left]
		idx++

		if right-left > 1 {
			for _, b := range []byte(strconv.Itoa(right - left)) {
				chars[idx] = b
				idx++
			}
		}
	}

	return idx
}
