package main // https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	maxSum := -100001
	maxLvl := 0

	nextLvl := make([]*TreeNode, 0, 10000)
	nextLvl = append(nextLvl, root)

	lvl := 0
	for {
		lvl++

		sum := 0
		size := len(nextLvl)

		for i := range size {
			sum += nextLvl[i].Val
			if nextLvl[i].Left != nil {
				nextLvl = append(nextLvl, nextLvl[i].Left)
			}
			if nextLvl[i].Right != nil {
				nextLvl = append(nextLvl, nextLvl[i].Right)
			}
		}
		nextLvl = nextLvl[size:]

		if sum > maxSum {
			maxSum = sum
			maxLvl = lvl
		}

		if len(nextLvl) == 0 {
			break
		}
	}

	return maxLvl
}
