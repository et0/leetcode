package lm2415 // https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursively(left *TreeNode, right *TreeNode, lvl int) {
	if left == nil || right == nil {
		return
	}
	if lvl%2 != 0 {
		left.Val, right.Val = right.Val, left.Val
	}

	recursively(left.Left, right.Right, lvl+1)
	recursively(right.Left, left.Right, lvl+1)

}

func reverseOddLevels(root *TreeNode) *TreeNode {
	recursively(root.Left, root.Right, 1)

	return root
}

func Wrapper(root *TreeNode) *TreeNode {
	return reverseOddLevels(root)
}
