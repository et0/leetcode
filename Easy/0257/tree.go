package le0257 // https://leetcode.com/problems/binary-tree-paths/description/

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	output := make([]string, 0, 100)
	path := make([]string, 0, 100)
	pathMap := make(map[*TreeNode]int, 0)
	trees := make([]*TreeNode, 0, 100)

	for root != nil {
		if root.Left != nil {
			path = append(path, strconv.Itoa(root.Val))
			pathMap[root] = len(path)

			current := root.Left
			root.Left = nil
			trees = append(trees, root)
			root = current
		} else if root.Right != nil {
			path = append(path, strconv.Itoa(root.Val))

			root = root.Right
		} else {
			path = append(path, strconv.Itoa(root.Val))
			if _, ok := pathMap[root]; !ok {
				output = append(output, strings.Join(path, "->"))
			}

			size := len(trees)
			if size == 0 {
				break
			}

			root = trees[size-1]
			trees = trees[:size-1]
			path = path[:pathMap[root]-1]
		}
	}

	return output
}

func Wrapper(root *TreeNode) []string {
	return binaryTreePaths(root)
}
