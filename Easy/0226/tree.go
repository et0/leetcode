package le0226 // https://leetcode.com/problems/invert-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Left != nil || root.Right != nil {
		root.Right, root.Left = root.Left, root.Right
	}
	_, _ = invertTree(root.Left), invertTree(root.Right)

	return root
}

func Wrapper(root *TreeNode) *TreeNode {
	return invertTree(root)
}
