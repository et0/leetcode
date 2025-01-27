package convert_l0108

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTree(head *TreeNode) {
	fmt.Println(head.Val)
	if head.Left != nil {
		PrintTree(head.Left)
	}
	if head.Right != nil {
		PrintTree(head.Right)
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	half := len(nums) / 2

	head := TreeNode{nums[half], nil, nil}
	if len(nums) == 1 {
		return &head
	}
	head.Left = sortedArrayToBST(nums[:half])
	head.Right = sortedArrayToBST(nums[half+1:])

	return &head
}

func Wrapper(nums []int) *TreeNode {
	return sortedArrayToBST(nums)
}
