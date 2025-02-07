package l_e_0003_string

func lengthOfLongestSubstring(s string) int {
	maxr := 0
	data := make(map[rune]int)
	for k, v := range s {
		if index, ok := data[v]; ok {
			if len(data) > maxr {
				maxr = len(data)
			}
			data = make(map[rune]int)
			for i := index + 1; i <= k; i++ {
				data[rune(s[i])] = i
			}
		} else {
			data[v] = k
		}
	}

	if len(data) > maxr {
		return len(data)
	}
	return maxr
}

func Wrapper(s string) int {
	return lengthOfLongestSubstring(s)
}
