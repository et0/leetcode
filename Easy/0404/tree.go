package le0404 // https://leetcode.com/problems/sum-of-left-leaves/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	prev := make([]*TreeNode, 0, 1000)
	for {
		if root.Left != nil {
			if root.Left.Left == nil && root.Left.Right == nil {
				sum += root.Left.Val
			}

			current := root.Left
			root.Left = nil
			prev = append(prev, root)
			root = current
		} else if root.Right != nil {
			root = root.Right
		} else {
			size := len(prev)
			if size == 0 {
				break
			}
			root = prev[size-1]
			prev = prev[:size-1]
		}
	}

	return sum
}

func Wrapper(root *TreeNode) int {
	return sumOfLeftLeaves(root)
}
