package main

import "testing"

type data struct {
	nums   []int
	result bool
}

func TestIsArraySpecial(t *testing.T) {
	tests := []data{
		{[]int{1}, true},
		{[]int{2, 1, 4}, true},
		{[]int{4, 3, 1, 6}, false},
	}

	for _, v := range tests {
		result := isArraySpecial(v.nums)
		if result != v.result {
			t.Error("Expected ", v.result, ", got ", result)
		}
	}
}
