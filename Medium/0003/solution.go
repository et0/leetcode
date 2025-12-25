package main // https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	uniq := make(map[byte]int)
	maxLength := 0

	for left, right := 0, 0; left < len(s); left++ {
		for ; right < len(s) && uniq[s[right]] == 0; right++ {
			uniq[s[right]]++
		}

		if right-left > maxLength {
			maxLength = right - left
		}

		uniq[s[left]]--
	}

	return maxLength
}
