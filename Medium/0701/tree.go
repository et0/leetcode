package lm0701 // https://leetcode.com/problems/insert-into-a-binary-search-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	head := root

	new := &TreeNode{val, nil, nil}
	if root == nil {
		return new
	}

	for {
		if root.Left != nil && root.Val > val {
			root = root.Left
		} else if root.Right != nil && root.Val < val {
			root = root.Right
		} else if root.Val > val {
			root.Left = new
			break
		} else {
			root.Right = new
			break
		}
	}

	return head
}

func Wrapper(root *TreeNode, val int) *TreeNode {
	return insertIntoBST(root, val)
}
