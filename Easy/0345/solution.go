package le0345 // https://leetcode.com/problems/reverse-vowels-of-a-string/description/

func reverseVowels(s string) string {
	newS := []byte(s)
	for left, right := 0, len(s)-1; left < right; left++ {
		if !isVowel(s[left]) {
			continue
		}

		for ; left <= right && !isVowel(s[right]); right-- {
		}

		newS[left], newS[right] = newS[right], newS[left]
		right--
	}

	return string(newS)
}

func isVowel(s byte) bool {
	switch s {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}

	return false
}

func Wrapper(s string) string {
	return reverseVowels(s)
}
