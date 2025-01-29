package pathsum

import (
	"testing"
)

type data struct {
	root      *TreeNode
	targetSum int
	result    bool
}

func TestWrapper(t *testing.T) {
	tests := []data{
		{
			&TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val:   13,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:  4,
						Left: nil,
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
			}, 18, true},
		{
			&TreeNode{
				Val:  -2,
				Left: nil,
				Right: &TreeNode{
					Val:   -3,
					Left:  nil,
					Right: nil,
				},
			}, -5, true},
		{
			&TreeNode{
				Val:  -2,
				Left: nil,
				Right: &TreeNode{
					Val:   -3,
					Left:  nil,
					Right: nil,
				},
			}, -5, true},
		{nil, 0, false},
	}

	for _, v := range tests {
		result := Wrapper(v.root, v.targetSum)
		if result != v.result {
			t.Error("Expected ", v.result, ", got ", result, ", for ", v.root)
		}
	}
}
