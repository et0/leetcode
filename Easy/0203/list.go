package le0203 // https://leetcode.com/problems/remove-linked-list-elements/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElementsIter(head *ListNode, val int) *ListNode {
	start := head
	var prev *ListNode

	for head != nil {
		if head.Val == val {
			if prev == nil {
				start = head.Next
			} else {
				prev.Next = head.Next
			}
		} else {
			prev = head
		}
		head = head.Next
	}

	return start
}

func removeElementsRecursiveBestRuntime(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return removeElementsRecursiveBestRuntime(head.Next, val)
	}
	head.Next = removeElementsRecursiveBestRuntime(head.Next, val)
	return head
}

func removeElementsRecursiveMy(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	next := removeElementsRecursiveMy(head.Next, val)
	if head.Val == val {
		return next
	}
	head.Next = next

	return head
}

func Wrapper(head *ListNode, val int) *ListNode {
	return removeElementsRecursiveMy(head, val)
}
