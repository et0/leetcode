package le0938 // https://leetcode.com/problems/range-sum-of-bst/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	stack := make([]*TreeNode, 0, 100000)
	sum := 0
	for {
		if root.Left != nil {
			current := root.Left
			root.Left = nil
			stack = append(stack, root)
			root = current
			continue
		}

		if root.Val >= low && root.Val <= high {
			sum += root.Val
		}

		if root.Right != nil {
			root = root.Right
			continue
		}

		size := len(stack)
		if size == 0 {
			break
		}
		root = stack[size-1]
		stack = stack[:size-1]
	}

	return sum
}

func Wrapper(root *TreeNode, low int, high int) int {
	return rangeSumBST(root, low, high)
}
