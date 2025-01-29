package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	} else if root.Left == nil && root.Right == nil && targetSum-root.Val == 0 {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func Wrapper(root *TreeNode, targetSum int) bool {
	return hasPathSum(root, targetSum)
}
