package lm2196 // https://leetcode.com/problems/create-binary-tree-from-descriptions/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type el struct {
	addr *TreeNode
	lvl  int
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	size := len(descriptions)
	hash := make(map[int]*el, size)

	for _, v := range descriptions {
		if _, ok := hash[v[0]]; !ok {
			hash[v[0]] = &el{addr: &TreeNode{Val: v[0]}}
		}
		if _, ok := hash[v[1]]; !ok {
			hash[v[1]] = &el{addr: &TreeNode{Val: v[1]}, lvl: 1}
		} else {
			hash[v[1]].lvl++
		}
		if v[2] == 1 {
			hash[v[0]].addr.Left = hash[v[1]].addr
		} else {
			hash[v[0]].addr.Right = hash[v[1]].addr
		}
	}

	for _, v := range hash {
		if v.lvl == 0 {
			return v.addr
		}
	}

	return nil
}

func Wrapper(descriptions [][]int) *TreeNode {
	return createBinaryTree(descriptions)
}
