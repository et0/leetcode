package main

import "fmt"

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

func main() {
	l1 := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}
	l2 := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

	result := addTwoNumbers(&l1, &l2)

	if result != nil {
		for {
			fmt.Println(result.Val)
			if result.Next == nil {
				break
			}

			result = result.Next
		}
	}
}
