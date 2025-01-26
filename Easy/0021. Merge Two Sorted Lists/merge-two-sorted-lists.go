package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := ListNode{0, nil}
	curr := &head

	for {
		if list1 == nil || list2 == nil {
			break
		}

		if list1.Val >= list2.Val {
			curr.Next = list2
			list2 = list2.Next
		} else {
			curr.Next = list1
			list1 = list1.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	}

	return head.Next
}

func main() {
	// 1,2,4
	list1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	// 1,3,4
	list2 := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	result := mergeTwoLists(&list1, &list2)

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
