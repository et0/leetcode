package romantointeger

func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	size := len(s)
	for i := 0; i < size; i++ {
		if i+1 < size && roman[s[i]] < roman[s[i+1]] {
			result -= roman[s[i]]
		} else {
			result += roman[s[i]]
		}
	}
	return result
}

func Wrapper(s string) int {
	return romanToInt(s)
}
