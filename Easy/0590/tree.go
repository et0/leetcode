package le0590 // https://leetcode.com/problems/n-ary-tree-postorder-traversal/description/

type Node struct {
	Val      int
	Children []*Node
}

func recursive(root *Node, output *[]int) {
	for _, c := range root.Children {
		recursive(c, output)
	}
	*output = append(*output, root.Val)
}

func postorder(root *Node) []int {
	output := make([]int, 0, 10000)
	if root == nil {
		return output
	}

	recursive(root, &output)

	return output
}

// iteratively
func postorderIteratively(root *Node) []int {
	output := make([]int, 0, 10000)
	if root == nil {
		return output
	}

	stack := make([]*Node, 0, 10000)
	for {
		if len(root.Children) == 0 {
			output = append(output, root.Val)

			size := len(stack)
			if size == 0 {
				break
			}

			root = stack[size-1]
			stack = stack[:size-1]

			continue
		}

		current := root.Children[0]
		root.Children = root.Children[1:]
		stack = append(stack, root)
		root = current
	}

	return output
}

func Wrapper(root *Node) []int {
	return postorder(root)
}
