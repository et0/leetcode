package le0222 // https://leetcode.com/problems/count-complete-tree-nodes/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func Wrapper(root *TreeNode) int {
	return countNodes(root)
}
