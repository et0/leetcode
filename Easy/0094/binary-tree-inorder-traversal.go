package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0, 100)
	values := make([]int, 0, 100)

	if root == nil {
		return make([]int, 0, 0)
	}

	for {
		fmt.Println(root, len(stack), stack)

		if root.Left != nil {
			current := *root
			current.Left = nil
			stack = append(stack, &current)

			root = root.Left
			continue
		}
		values = append(values, root.Val)

		if root.Right != nil {
			root = root.Right
			continue
		}

		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			continue
		}

		break
	}

	return values
}

func main() {
	tree := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
	}

	fmt.Println(inorderTraversal(&tree))
}
