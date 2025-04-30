package lh1028 // https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/description/

import (
	"regexp"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(traversal string) *TreeNode {
	stack := make([]*TreeNode, 0, 1000)

	re, _ := regexp.Compile(`(\-*)(\d+)`)
	find := re.FindAllStringSubmatch(traversal, -1)

	val, _ := strconv.Atoi(find[0][2])
	root := &TreeNode{Val: val}

	head := root

	for i := 1; i < len(find); i++ {
		lvl := len(find[i][1])
		val, _ = strconv.Atoi(find[i][2])

		if lvl > len(stack) {
			stack = append(stack, root)
			root.Left = &TreeNode{Val: val}
			root = root.Left
		} else {
			root = stack[lvl-1]
			stack = stack[:lvl]
			root.Right = &TreeNode{Val: val}
			root = root.Right
		}
	}

	return head
}

func Wrapper(traversal string) *TreeNode {
	return recoverFromPreorder(traversal)
}
