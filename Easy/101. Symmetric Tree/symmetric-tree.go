package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func iteratively(root *TreeNode) bool {
	stackLeft, stackRight := make([]*TreeNode, 0, 1000), make([]*TreeNode, 0, 1000)
	currentLeft, currentRight := root.Left, root.Right

	for {
		if currentLeft == nil && currentRight == nil {
			if len(stackLeft) == 0 && len(stackRight) == 0 {
				return true
			}
			return false
		} else if currentLeft == nil || currentRight == nil || currentLeft.Val != currentRight.Val {
			return false
		}

		if currentLeft.Left != nil {
			current := *currentLeft
			current.Left = nil
			stackLeft = append(stackLeft, &current)

			currentLeft = currentLeft.Left
		} else if currentLeft.Right != nil {
			currentLeft = currentLeft.Right
		} else if len(stackLeft) > 0 {
			currentLeft = stackLeft[len(stackLeft)-1]
			stackLeft = stackLeft[:len(stackLeft)-1]
		} else {
			currentLeft = nil
		}

		if currentRight.Right != nil {
			current := *currentRight
			current.Right = nil
			stackRight = append(stackRight, &current)
			currentRight = currentRight.Right
		} else if currentRight.Left != nil {
			currentRight = currentRight.Left
		} else if len(stackRight) > 0 {
			currentRight = stackRight[len(stackRight)-1]
			stackRight = stackRight[:len(stackRight)-1]
		} else {
			currentRight = nil
		}
	}
}

func recursively(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return recursively(left.Left, right.Right) && recursively(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return iteratively(root)
	// return recursively(root.Left, root.Right)
}

func main() {
	tree := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(isSymmetric(&tree))
}
