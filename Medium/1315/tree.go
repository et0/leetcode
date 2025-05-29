package lm1315 // https://leetcode.com/problems/sum-of-nodes-with-even-valued-grandparent/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Val%2 == 0 {
		if root.Left != nil {
			if root.Left.Left != nil {
				sum += root.Left.Left.Val
			}
			if root.Left.Right != nil {
				sum += root.Left.Right.Val
			}
		}
		if root.Right != nil {
			if root.Right.Left != nil {
				sum += root.Right.Left.Val
			}
			if root.Right.Right != nil {
				sum += root.Right.Right.Val
			}
		}
	}

	if root.Left != nil {
		sum += sumEvenGrandparent(root.Left)
	}
	if root.Right != nil {
		sum += sumEvenGrandparent(root.Right)
	}

	return sum
}

func Wrapper(root *TreeNode) int {
	return sumEvenGrandparent(root)
}
