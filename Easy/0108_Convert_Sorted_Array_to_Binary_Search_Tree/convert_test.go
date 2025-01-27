package convert_l0108

import (
	"testing"
)

type data struct {
	nums   []int
	result []int
}

func Tree2Slice(head *TreeNode, slice *[]int) {
	*slice = append(*slice, head.Val)

	if head.Left != nil {
		Tree2Slice(head.Left, slice)
	}
	if head.Right != nil {
		Tree2Slice(head.Right, slice)
	}
}

func TestWrapper(t *testing.T) {
	tests := []data{
		{[]int{-10, -3, 0, 5, 9}, []int{0, -3, -10, 9, 5}},
		{[]int{1, 3}, []int{3, 1}},
	}

	for _, v := range tests {
		result := make([]int, 0, 10000)
		Tree2Slice(Wrapper(v.nums), &result)

		if len(v.result) != len(result) {
			t.Error("Expected ", v.result, ", got ", result)
		}

		for k, vv := range result {
			if v.result[k] != vv {
				t.Error("Expected ", v.result, ", got ", result)
			}
		}
	}
}
