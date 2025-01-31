package balancedbinarytree

import "testing"

type data struct {
	root   *TreeNode
	result bool
}

func TestWrapper(t *testing.T) {
	tests := []data{
		{
			&TreeNode{Val: 6, Left: &TreeNode{Val: 20, Left: &TreeNode{Val: 30, Left: &TreeNode{Val: 40, Left: nil, Right: nil}, Right: &TreeNode{Val: 41, Left: nil, Right: &TreeNode{Val: 51, Left: &TreeNode{Val: 60, Left: nil, Right: nil}, Right: &TreeNode{Val: 61, Left: nil, Right: nil}}}}, Right: nil}, Right: &TreeNode{Val: 21, Left: &TreeNode{Val: 30, Left: nil, Right: nil}, Right: &TreeNode{Val: 31, Left: &TreeNode{Val: 40, Left: nil, Right: nil}, Right: nil}}},
			false,
		},
		{
			nil,
			true,
		},
		{
			&TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: nil, Right: nil}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}}},
			true,
		},
	}

	for _, v := range tests {
		result := Wrapper(v.root)
		if result != v.result {
			t.Error("Expected ", v.result, ", got ", result, ", for ", v.root)
		}
	}
}
