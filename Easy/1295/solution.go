package main // https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/

func findNumbers(nums []int) int {
	even := 0
	for _, n := range nums {
		if (n > 9 && n < 100) || (n > 999 && n < 10000) || n == 100000 {
			even++
		}
	}
	return even
}
