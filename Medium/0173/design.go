package lm0173 // https://leetcode.com/problems/binary-search-tree-iterator/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	out := BSTIterator{stack: make([]*TreeNode, 0, 100000)}
	out.push(root)

	return out
}

func (b *BSTIterator) push(next *TreeNode) {
	for next != nil {
		b.stack = append(b.stack, next)
		next = next.Left
	}
}
func (b *BSTIterator) Next() int {
	node := b.stack[len(b.stack)-1]
	b.stack = b.stack[:len(b.stack)-1]

	b.push(node.Right)

	return node.Val
}

func (b *BSTIterator) HasNext() bool {
	return len(b.stack) > 0
}
