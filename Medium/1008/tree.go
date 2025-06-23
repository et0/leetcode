package lm1008 // https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}
	head := root

	stack := make([]*TreeNode, 0, 100)

	size := len(preorder)
	for i := 1; i < size; i++ {
		stack = append(stack, root)

		if preorder[i] < root.Val {
			root.Left = &TreeNode{Val: preorder[i]}
			root = root.Left

			continue
		}

		sizeStack := len(stack)
		for j := 0; j < sizeStack; j++ {
			if preorder[i] < stack[j].Val {
				continue
			}

			stack[j].Right = &TreeNode{Val: preorder[i]}
			root = stack[j].Right
			stack[j] = root
			stack = stack[:j+1]

			break
		}

	}

	return head
}

func Wrapper(preorder []int) *TreeNode {
	return bstFromPreorder(preorder)
}
