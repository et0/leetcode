package le0206 // https://leetcode.com/problems/reverse-linked-list/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

// Iteratively
func reverseListIteratively(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev *ListNode
	for {
		current := head.Next
		head.Next = prev
		prev = head

		if current == nil {
			break
		}
		head = current
	}

	return head
}

func reverseListRecursive(head *ListNode) *ListNode {
	revHead := head
	if head != nil && head.Next != nil {
		revHead = reverseListRecursive(head.Next)
		head.Next.Next = head
		head.Next = nil
	}

	return revHead
}

func Wrapper(head *ListNode) *ListNode {
	return reverseListRecursive(head)
}
