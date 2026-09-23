package main // https://leetcode.com/problems/group-anagrams/

import (
	"slices"
	"strconv"
	"strings"
)

// groupAnagramsAnother группирует строки-анаграммы вместе.
// Ключ группы — 26 счётчиков букв (a-z), склеенных в строку через разделитель.
// Анаграммы дают одинаковый ключ. O(n*k), где n — число строк, k — длина строки:
// на строку уходит один проход, а сборка ключа из 26 чисел — константа.
// В Go ключом map может быть и сам массив [26]int: массивы фиксированного размера
// сравнимы, поэтому склеивать строку не обязательно. Слайс []int ключом быть не может.
func groupAnagramsAnother(strs []string) [][]string {
	// preOut: ключ (сигнатура анаграммы) -> список строк с этой сигнатурой
	preOut := make(map[string][]string, len(strs))

	for _, str := range strs {
		// hash[i] = сколько раз буква ('a'+i) встречается в str
		hash := [26]int{}

		for _, b := range str {
			hash[b-97]++ // 'a' = 97, поэтому b-97 даёт индекс 0..25
		}

		// Строим ключ: "1_0_0_..._2_" — одинаков для всех анаграмм
		// Разделитель "_" нужен, чтобы [1,11] и [11,1] не склеились в "111"
		var sb strings.Builder
		for _, v := range hash {
			sb.WriteString(strconv.Itoa(v) + "_")
		}
		key := sb.String()

		// Добавляем строку в группу с этим ключом
		preOut[key] = append(preOut[key], str)
	}

	out := make([][]string, 0, len(preOut))
	for _, v := range preOut {
		out = append(out, v)
	}

	return out
}

// groupAnagrams группирует строки-анаграммы вместе.
// Ключ группы — отсортированная строка: анаграммы дают одинаковый ключ
// ("eat", "tea", "ate" -> "aet"). O(n * k log k), где n — число строк, k — длина строки.
// k log k даёт сортировка ключа. Операции с map по строковому ключу тоже не бесплатны:
// хеш считается по всем байтам, то есть O(k), и на общий порядок это не влияет.
func groupAnagrams(strs []string) [][]string {
	// groups: ключ (отсортированная строка) -> список анаграмм
	groups := make(map[string][]string, len(strs))

	for _, str := range strs {
		// Превращаем строку в слайс байт, чтобы отсортировать in-place
		byteStr := []byte(str)
		slices.Sort(byteStr)

		// Отсортированная строка — общий ключ для всех анаграмм
		key := string(byteStr)
		groups[key] = append(groups[key], str)
	}

	// Преобразуем map в слайс слайсов (порядок групп не важен)
	out := make([][]string, 0, len(groups))
	for _, group := range groups {
		out = append(out, group)
	}

	return out
}
