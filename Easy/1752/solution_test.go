package l_e_1752_array

import "testing"

type data struct {
	nums   []int
	result bool
}

func TestWrapper(t *testing.T) {
	tests := []data{
		{[]int{1, 1, 1}, true},
		{[]int{}, true},
		{[]int{4, 3, 1, 6}, false},
		{[]int{3, 4, 5, 1, 2}, true},
		{[]int{2, 1, 3, 4}, false},
		{[]int{1, 2, 3}, true},
	}

	for _, v := range tests {
		result := Wrapper(v.nums)
		if result != v.result {
			t.Error("Expected ", v.result, ", got ", result, " for ", v.nums)
		}
	}
}
