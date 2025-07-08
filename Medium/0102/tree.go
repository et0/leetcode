package lm0102 // https://leetcode.com/problems/binary-tree-level-order-traversal/

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

	(*out)[lvl] = append((*out)[lvl], root.Val)

	recursion(root.Left, out, lvl+1)
	recursion(root.Right, out, lvl+1)
}

func levelOrder(root *TreeNode) [][]int {
	out := make([][]int, 0, 2000)

	recursion(root, &out, 0)

	return out
}

func Wrapper(root *TreeNode) [][]int {
	return levelOrder(root)
}
