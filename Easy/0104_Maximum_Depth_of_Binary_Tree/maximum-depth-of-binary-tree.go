package maximumdepthofbinarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	depth := struct {
		current int
		max     int
	}{0, 0}

	type data struct {
		lvl  int
		tree *TreeNode
	}
	stack := make([]data, 0, 10000)

	if root == nil {
		return depth.max
	}
	depth.current++

	// Прямой обход дерева
	for {
		fmt.Println(depth)
		if root.Left != nil {
			current := *root
			current.Left = nil
			stack = append(stack, data{depth.current, &current})

			root = root.Left
			depth.current++

			continue
		}
		if root.Right != nil {
			root = root.Right
			depth.current++

			continue
		}

		if depth.current > depth.max {
			depth.max = depth.current
		}
		if len(stack) > 0 {
			// set root
			root = stack[len(stack)-1].tree
			// set lvl
			depth.current = stack[len(stack)-1].lvl
			// pop item from stack
			stack = stack[:len(stack)-1]

			continue
		}

		break
	}

	return depth.max
}
