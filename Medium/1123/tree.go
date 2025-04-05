package lm1123 // https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, lca := recursion(root)
	return lca
}

func recursion(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	leftDepth, leftLCA := recursion(root.Left)
	rightDepth, rightLCA := recursion(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1, leftLCA
	}
	if leftDepth < rightDepth {
		return rightDepth + 1, rightLCA
	}
	return leftDepth + 1, root
}

func Wrapper(root *TreeNode) *TreeNode {
	return lcaDeepestLeaves(root)
}
