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
}

func Wrapper(root *TreeNode) bool {
	return checkTree(root)
}
