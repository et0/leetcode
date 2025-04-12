package le0205 // https://leetcode.com/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	direct := make(map[rune]rune)
	reverse := make(map[rune]rune)

	for k, v := range s {
		if _, ok := direct[v]; !ok {
			direct[v] = rune(t[k])
		} else if direct[v] != rune(t[k]) {
			return false
		}

		if _, ok := reverse[rune(t[k])]; !ok {
			reverse[rune(t[k])] = v
		} else if reverse[rune(t[k])] != v {
			return false
		}
	}

	return true
}

func Wrapper(s string, t string) bool {
	return isIsomorphic(s, t)
}
