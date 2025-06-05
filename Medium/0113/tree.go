package lm0113 // https://leetcode.com/problems/path-sum-ii/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursion(root *TreeNode, targetSum int, vals *[][]int, subVal *[]int) {
	*subVal = append(*subVal, root.Val)

	if root.Left == nil && root.Right == nil {
		if targetSum-root.Val == 0 {
			*vals = append(*vals, *subVal)
		}

		return
	}

	if root.Left != nil {
		newSubVal := make([]int, 0, 5000)
		newSubVal = append(newSubVal, *subVal...)

		recursion(root.Left, targetSum-root.Val, vals, &newSubVal)
	}

	if root.Right != nil {
		newSubVal := make([]int, 0, 5000)
		newSubVal = append(newSubVal, *subVal...)

		recursion(root.Right, targetSum-root.Val, vals, &newSubVal)
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	vals := make([][]int, 0, 5000)

	recursion(root, targetSum, &vals, &[]int{})

	return vals
}

func Wrapper(root *TreeNode, targetSum int) [][]int {
	return pathSum(root, targetSum)
}
