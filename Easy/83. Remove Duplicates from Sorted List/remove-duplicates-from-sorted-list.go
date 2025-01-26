package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	first := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}

	return first
}

func main() {
	// 1,2,4
	list := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}

	result := deleteDuplicates(&list)

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
