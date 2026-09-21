package main // https://leetcode.com/problems/isomorphic-strings/

// Идея: соответствие символов должно быть взаимно однозначным — символ s
// переходит ровно в один символ t, и символ t занят ровно одним символом s.
// Поэтому две map: sToT (прямое) и tToS (обратное), на каждой позиции проверяем обе.
// Время: O(n) — один проход, на каждом шаге константное число операций с map,
// каждая O(1) в среднем.
// Память: O(1) — в каждой map не больше 128 ключей (алфавит ASCII),
// от длины строки размер не зависит.
// Ловушка: одной проверки из s в t мало — она пропускает случай, когда два разных
// символа s переходят в один символ t ("ab" / "cc").
func isIsomorphic(s string, t string) bool {
	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if v, exists := sToT[s[i]]; exists && v != t[i] {
			return false
		}

		if v, exists := tToS[t[i]]; exists && v != s[i] {
			return false
		}

		sToT[s[i]] = t[i]
		tToS[t[i]] = s[i]
	}

	return true
}
