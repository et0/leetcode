package l_e_0111_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	} else if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	} else if root.Left == nil && root.Right == nil {
		return 1
	}

	return min(minDepth(root.Left)+1, minDepth(root.Right)+1)
}

func Wrapper(root *TreeNode) int {
	return minDepth(root)
}
