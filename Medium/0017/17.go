package lm0017 // https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phone := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	size := 1 // size of combination
	for _, v := range digits {
		size *= len(phone[v])
	}

	result := make([]string, size)
	repeat := size
	for _, v := range digits {
		index := 0
		repeat /= len(phone[v])
		for i := 0; i < size; i++ {
			result[i] += string(phone[v][index])
			if (i+1)%repeat == 0 {
				index++
				if index == len(phone[v]) {
					index = 0
				}
			}
		}
	}

	return result
}

func Wrapper(digits string) []string {
	return letterCombinations(digits)
}
