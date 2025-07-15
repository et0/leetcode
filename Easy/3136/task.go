package le3136 // https://leetcode.com/problems/valid-word/description/

import "unicode"

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	var (
		digit, upper, lower, vowel, consonant bool
	)
	for _, w := range word {
		if unicode.IsDigit(w) {
			digit = true

			continue
		}

		if !unicode.IsLetter(w) {
			return false
		}

		if unicode.IsUpper(w) {
			upper = true
			w = unicode.ToLower(w)
		} else {
			lower = true
		}

		if w == 'a' || w == 'e' || w == 'i' || w == 'o' || w == 'u' {
			vowel = true
		} else {
			consonant = true
		}

	}

	return (digit || upper || lower) && vowel && consonant
}

func Wrapper(word string) bool {
	return isValid(word)
}
