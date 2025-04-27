package le2331 // https://leetcode.com/problems/evaluate-boolean-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	switch root.Val {
	case 1:
		return true
	case 2:
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	case 3:
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	default:
		return false
	}
}

func Wrapper(root *TreeNode) bool {
	return evaluateTree(root)
}
