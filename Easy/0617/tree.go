package le0617 // https://leetcode.com/problems/merge-two-binary-trees/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type data struct {
	root1 *TreeNode
	root2 *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	}

	stack := make([]data, 0, 2000)
	head := root1

	for {
		if root1 != nil && root2 != nil {
			root1.Val += root2.Val

			if root1.Left != nil {
				if root2.Left != nil {
					if root1.Right == nil {
						root1.Right = root2.Right
					} else if root2.Right == nil {
						root2.Right = root1.Right
					} else {
						stack = append(stack, data{root1.Right, root2.Right})
					}

					root1, root2 = root1.Left, root2.Left
					continue
				}
			} else {
				root1.Left = root2.Left
			}

			if root1.Right != nil {
				if root2.Right != nil {
					root1, root2 = root1.Right, root2.Right

					continue
				}
			} else {
				root1.Right = root2.Right
			}
		}

		size := len(stack)
		if size == 0 {
			break
		}
		root1, root2 = stack[size-1].root1, stack[size-1].root2
		stack = stack[:size-1]
	}

	return head
}

func Wrapper(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return mergeTrees(root1, root2)
}
