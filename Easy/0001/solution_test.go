package main

import "testing"

type data struct {
	nums   []int
	target int
	result []int
}

func TestTwoSum(t *testing.T) {
	tests := []data{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, v := range tests {
		result := twoSum(v.nums, v.target)
		if result[0] != v.result[0] || result[1] != v.result[1] {
			t.Error("Expected ", v.result, ", got ", result)
		}
	}
}
