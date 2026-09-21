package main // https://leetcode.com/problems/compare-version-numbers/

import (
	"strconv"
)

// search извлекает числовой компонент версии, начиная с позиции left.
// Возвращает значение компонента и позицию следующего компонента (после точки).
// Пустой компонент (конец строки или сразу точка) -> 0.
func search(left int, str string) (int, int) {
	right := left
	for ; right < len(str) && str[right] != '.'; right++ {

	}

	// Пустой компонент (например, "1..2" или конец строки) -> считаем как 0
	if right == left {
		return 0, right + 1
	}

	number, _ := strconv.Atoi(str[left:right])

	return number, right + 1
}

// compareVersion сравнивает две версии, разбитые точками на числовые компоненты.
// Недостающие компоненты считаются нулями. Ведущие нули игнорируются.
// Возвращает: -1 если version1 < version2, 0 если равны, 1 если version1 > version2.
// O(n + m) времени, O(1) памяти
func compareVersion(version1 string, version2 string) int {
	left1, left2 := 0, 0
	for left1 < len(version1) || left2 < len(version2) {
		var number1, number2 int

		number1, left1 = search(left1, version1)
		number2, left2 = search(left2, version2)

		if number1 > number2 {
			return 1
		} else if number1 < number2 {
			return -1
		}

	}

	return 0
}
