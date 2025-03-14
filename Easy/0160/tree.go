package le0160 // https://leetcode.com/problems/intersection-of-two-linked-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := make(map[*ListNode]bool), make(map[*ListNode]bool)

	for headA != nil || headB != nil {
		if headA != nil {
			a[headA] = true
			if b[headA] {
				return headA
			}
			headA = headA.Next
		}
		if headB != nil {
			b[headB] = true
			if a[headB] {
				return headB
			}
			headB = headB.Next
		}
	}

	return nil
}

func Wrapper(headA, headB *ListNode) *ListNode {
	return getIntersectionNode(headA, headB)
}
