package convert_l0108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
