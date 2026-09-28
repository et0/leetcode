package main // https://leetcode.com/problems/valid-palindrome/

func isValidSymbol(s byte) bool {
	return (s >= '0' && s <= '9') || (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}

func toLower(s byte) byte {
	if s >= 'A' && s <= 'Z' {
		return s + 32
	}

	return s
}

// Идея: два указателя идут с концов строки навстречу друг другу. Символы, которые не буквы
// и не цифры, просто пропускаются — строку не чистим и не копируем. Буквы сравниваем,
// приведя обе к нижнему регистру прямо в момент сравнения.
// Время: O(n) — каждый указатель проходит свою часть строки, вместе они проходят её один раз.
// Память: O(1) — ни новых строк, ни срезов, только два индекса. Вариант, где строку сначала
// чистят в новый срез, дал бы O(n).
// Ловушка: сдвиг на 32 переводит регистр только у букв. Если применить его ко всем символам,
// цифра совпадёт с буквой: '1'+32 это 'Q', а '0'+32 это 'P'. Поэтому в toLower стоит проверка
// диапазона, а тест "0P" ловит её отсутствие.
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isValidSymbol(s[left]) {
			left++
			continue
		}

		if !isValidSymbol(s[right]) {
			right--
			continue
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}
