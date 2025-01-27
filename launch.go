package main

// twosum "github.com/et0/leetcode/Easy/0001_two_sum"
import (
	l0021 "github.com/et0/leetcode/Easy/0021_Merge_Two_Sorted_Lists"
)

func main() {
	l0021.Wrapper(&l0021.ListNode{Val: 1, Next: &l0021.ListNode{Val: 2, Next: &l0021.ListNode{Val: 4, Next: nil}}}, &l0021.ListNode{Val: 1, Next: &l0021.ListNode{Val: 3, Next: &l0021.ListNode{Val: 4, Next: nil}}})
	// fmt.Println(validparentheses.Wrapper("()({[]})[]")) 0020
	// fmt.Println(longestcommonprefix.Wrapper([]string{"a"})) // 0014
	// fmt.Println(romantointeger.Wrapper("MCMXCIV")) // 0013
	// fmt.Println(palindromenumber.Wrapper(6466)) // 0009
	// fmt.Println(twosum.Wrapper([]int{1, 2, 3, 4}, 100)) // 0001
}
