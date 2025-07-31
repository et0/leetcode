package lm0814 // https://leetcode.com/problems/binary-tree-pruning/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursion(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := recursion(root.Left)
	if left == 0 {
		root.Left = nil
	}

	right := recursion(root.Right)
	if right == 0 {
		root.Right = nil
	}

	return root.Val + left + right
}

func pruneTree(root *TreeNode) *TreeNode {
	if recursion(root) == 0 {
		return nil
	}

	return root
}

func Wrapper(root *TreeNode) *TreeNode {
	return pruneTree(root)
}
