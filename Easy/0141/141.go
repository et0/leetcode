package le0141 // https://leetcode.com/problems/linked-list-cycle/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

// Two Points
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func hasCycle_Hash(head *ListNode) bool {
	hash := make(map[*ListNode]bool)

	for head != nil {
		if hash[head] {
			return true
		}

		hash[head] = true

		head = head.Next
	}

	return false
}

func hasCycle_SetMark(head *ListNode) bool {
	for head != nil {
		if head.Val == -1000000 {
			return true
		}

		head.Val = -1000000
		head = head.Next
	}

	return false
}

func Wrapper(head *ListNode) bool {
	return hasCycle(head)
}
