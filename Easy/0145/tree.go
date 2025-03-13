package le0145 // postorder traversal https://leetcode.com/problems/binary-tree-postorder-traversal/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	numbers := make([]int, 0, 100)
	trees := make([]*TreeNode, 0, 100)

	for root != nil {
		if root.Left != nil {
			current := root.Left
			root.Left = nil
			trees = append(trees, root)
			root = current
		} else if root.Right != nil {
			current := root.Right
			root.Right = nil
			trees = append(trees, root)
			root = current
		} else {
			numbers = append(numbers, root.Val)

			size := len(trees)
			if size > 0 {
				root = trees[size-1]
				trees = trees[:size-1]
			} else {
				root = nil
			}
		}
	}

	return numbers
}

func Wrapper(root *TreeNode) []int {
	return postorderTraversal(root)
}
