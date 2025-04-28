package lm1038 // https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recurs(root *TreeNode, sum *int) {
	if root.Right != nil {
		recurs(root.Right, sum)
	}

	root.Val += *sum
	*sum = root.Val

	if root.Left != nil {
		recurs(root.Left, sum)
	}
}

func bstToGst(root *TreeNode) *TreeNode {
	head, sum := root, 0
	recurs(root, &sum)

	return head
}

func bstToGstDefault(root *TreeNode) *TreeNode {
	sum := 0
	head := root
	stack := make([]*TreeNode, 0, 100)
	saved := make(map[*TreeNode]bool, 100)
	for {
		if root.Right != nil && !saved[root.Right] {
			saved[root.Right] = true
			stack = append(stack, root)
			root = root.Right
		} else {
			root.Val += sum
			sum = root.Val

			if root.Left != nil {
				root = root.Left
				continue
			}

			size := len(stack)
			if size > 0 {
				root = stack[size-1]
				stack = stack[:size-1]
			} else {
				break
			}
		}
	}

	return head
}

func Wrapper(root *TreeNode) *TreeNode {
	return bstToGst(root)
}
