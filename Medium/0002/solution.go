package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0, nil}
	first := result

	for {
		if l1 == nil && l2 == nil {
			break
		}

		if l1 != nil {
			result.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			result.Val += l2.Val
			l2 = l2.Next
		}

		if result.Val > 9 {
			result.Val %= 10
			result.Next = &ListNode{Val: 1, Next: nil}
		} else if l1 != nil || l2 != nil {
			result.Next = &ListNode{Val: 0, Next: nil}
		}
		result = result.Next
	}

	return first
}
