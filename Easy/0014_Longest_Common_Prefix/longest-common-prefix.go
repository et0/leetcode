package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	min := 0
	size := len(strs)

	for i := 1; i < size; i++ {
		if len(strs[min]) > len(strs[i]) {
			min = i
		}
	}

	last := 0
	// перебор каждой буквы короткого слова
	for k, v := range strs[min] {
		// перебор оставщихся слов
		for i := 0; i < size; i++ {
			// если это то самое слово, которое мы выбрали самым коротким
			if i == min {
				continue
			}

			if strs[i][k] != byte(v) {
				return strs[min][0:k]
			}
		}
		last++
	}

	return strs[min][0:last]
}

func Wrapper(strs []string) string {
	return longestCommonPrefix(strs)
}
