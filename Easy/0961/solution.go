package main // https://leetcode.com/problems/n-repeated-element-in-size-2n-array/

func repeatedNTimes(nums []int) int {
	n := len(nums) / 2
	uniq := make(map[int]int, n+1)
	for _, v := range nums {
		uniq[v]++
	}
	for k, v := range uniq {
		if v == n {
			return k
		}
	}

	return 0
}
