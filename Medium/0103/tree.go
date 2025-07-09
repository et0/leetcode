package lm0103 // https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursion(root *TreeNode, out *[][]int, lvl int) {
	if root == nil {
		return
	}

	if len(*out) == lvl {
		(*out) = append((*out), make([]int, 0, 2000))
	}

	if lvl%2 == 0 {
		(*out)[lvl] = append((*out)[lvl], root.Val)
	} else {
		(*out)[lvl] = append([]int{root.Val}, (*out)[lvl]...)
	}

	recursion(root.Left, out, lvl+1)
	recursion(root.Right, out, lvl+1)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	out := make([][]int, 0, 2000)

	recursion(root, &out, 0)

	return out
}

func Wrapper(root *TreeNode) [][]int {
	return zigzagLevelOrder(root)
}
