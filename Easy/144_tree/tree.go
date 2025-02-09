package l_e_0144_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	values := make([]int, 0, 100)
	roots := make([]*TreeNode, 0, 100)
	flag := true // flag for save value
	for {
		if root == nil {
			break
		}

		if flag {
			values = append(values, root.Val)
		} else {
			flag = true
		}

		if root.Left != nil {
			current := root.Left
			root.Left = nil
			roots = append(roots, root)
			root = current
		} else if root.Right != nil {
			root = root.Right
		} else if len(roots) > 0 {
			flag = false
			root = roots[len(roots)-1]
			roots = roots[:len(roots)-1]
		} else {
			break
		}
	}

	return values
}

func Wrapper(root *TreeNode) []int {
	return preorderTraversal(root)
}
