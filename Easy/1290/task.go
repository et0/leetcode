package le1290 // https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func recursion(head *ListNode) (int, float64) {
	if head.Next == nil {
		return head.Val * int(math.Pow(2, 0)), 1
	}

	result, lvl := recursion(head.Next)

	return result + head.Val*int(math.Pow(2, lvl)), lvl + 1
}

func getDecimalValue(head *ListNode) int {
	out, _ := recursion(head)

	return out
}

func Wrapper(head *ListNode) int {
	return getDecimalValue(head)
}
