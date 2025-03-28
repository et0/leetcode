package le0234 // https://leetcode.com/problems/palindrome-linked-list/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// for case with [] or [123]
	if head == nil || head.Next == nil {
		return true
	}

	list := make([]int, 0, 50000)
	twoPoint := head

	for head.Next != nil {
		val := head.Val
		head = head.Next

		if twoPoint.Next == nil {
			break
		}
		list = append(list, val)
		if twoPoint.Next.Next == nil {
			break
		}
		twoPoint = twoPoint.Next.Next
	}

	for head != nil {
		size := len(list)
		if size == 0 || head.Val != list[size-1] {
			return false
		}
		list = list[:size-1]
		head = head.Next
	}

	return true
}

func Wrapper(head *ListNode) bool {
	return isPalindrome(head)
}
