package le3477 // https://leetcode.com/problems/fruits-into-baskets-ii/description/

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	countFruits := len(fruits)

	head := &Node{Val: baskets[len(baskets)-1]}
	for i := len(baskets) - 2; i >= 0; i-- {
		head = &Node{
			Val:  baskets[i],
			Next: head,
		}
		head.Next.Prev = head
	}

	for _, v := range fruits {
		cur := head
		for ; cur != nil; cur = cur.Next {
			if v > cur.Val {
				continue
			}

			countFruits--
			// change Prev, if cur not first
			if cur.Prev == nil {
				head = cur.Next
			} else {
				cur.Prev.Next = cur.Next
			}
			// change Next
			if cur.Next != nil {
				cur.Next.Prev = cur.Prev
			}

			break
		}
	}

	return countFruits
}

func Wrapper(fruits []int, baskets []int) int {
	return numOfUnplacedFruits(fruits, baskets)
}
