package lm0654 // https://leetcode.com/problems/maximum-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	size := len(nums)
	if size == 0 {
		return nil
	}

	max := 0
	for i := 1; i < size; i++ {
		if nums[i] > nums[max] {
			max = i
		}
	}

	root := &TreeNode{
		Val:   nums[max],
		Left:  constructMaximumBinaryTree(nums[:max]),
		Right: constructMaximumBinaryTree(nums[max+1:]),
	}

	return root
}

func Wrapper(nums []int) *TreeNode {
	return constructMaximumBinaryTree(nums)
}
