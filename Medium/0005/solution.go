package main // https://leetcode.com/problems/longest-palindromic-substring/description/

import "fmt"

func isPalindrom(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func longestPalindrome(s string) string {
	checked := make(map[string]bool)

	max := struct {
		len   int
		left  int
		right int
	}{1, 0, 0}

	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] != s[j] {
				continue
			}

			check, ok := checked[s[i:j+1]]
			if ok {

			} else if len(s[i:j+1]) > 3 {
				checkSub, ok := checked[s[i+1:j]]
				if !ok {
					checkSub = isPalindrom(s[i+1 : j])
					checked[s[i+1:j]] = checkSub
				}
				check = checkSub
				checked[s[i:j+1]] = checkSub
			} else {
				check = true
				checked[s[i:j+1]] = check
			}

			if check && max.len < len(s[i:j+1]) {
				max.len, max.left, max.right = len(s[i:j+1]), i, j
			}
		}
	}

	return s[max.left : max.right+1]
}

func main() {
	//
	fmt.Println(longestPalindrome("babad"))
}
