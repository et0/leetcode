package le0700

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	stack := make([]*TreeNode, 5000)

	for {
		if root == nil || root.Val == val {
			break
		}

		if root.Left != nil {
			cur := root.Left
			root.Left = nil
			stack = append(stack, root)
			root = cur
		} else if root.Right != nil {
			root = root.Right
		} else {
			size := len(stack)
			if size == 0 {
				break
			}

			root = stack[size-1]
			stack = stack[:size-1]
		}
	}

	return root
}

func Wrapper(root *TreeNode, val int) *TreeNode {
	return searchBST(root, val)
}
