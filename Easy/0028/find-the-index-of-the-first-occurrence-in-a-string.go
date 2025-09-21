package l0028

func strStr(haystack string, needle string) int {

	sizeNeedle, sizeHaystack := len(needle), len(haystack)

	for i := 0; i <= sizeHaystack-sizeNeedle; i++ {
		if haystack[i:i+sizeNeedle] == needle {
			return i
		}
	}

	return -1
}

func Wrapper(haystack string, needle string) int {
	return strStr(haystack, needle)
}
