package le0242 // https://leetcode.com/problems/valid-anagram/description/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	memory := make(map[rune]int)
	for _, v := range s {
		memory[v]++
	}

	for _, v := range t {
		if memory[v] == 0 {
			return false
		}
		memory[v]--
	}

	return true
}

func Wrapper(s string, t string) bool {
	return isAnagram(s, t)
}
