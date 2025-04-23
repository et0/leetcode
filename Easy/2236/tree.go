package le2236 // https://leetcode.com/problems/root-equals-sum-of-children/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return root.Left.Val+root.Right.Val == root.Val
}

func Wrapper(root *TreeNode) bool {
	return checkTree(root)
}
