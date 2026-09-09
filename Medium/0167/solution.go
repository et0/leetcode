package main // https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

import (
	"fmt"
)

// twoSum находит два числа в отсортированном массиве, дающих в сумме target.
// Используется метод двух указателей (сужение диапазона) за O(n).
// Возвращает индексы (1-based) в порядке возрастания.
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		if numbers[left]+numbers[right] == target {
			break
		}

		if numbers[left]+numbers[right] > target {
			right-- // сумма слишком большая - двигаем правый указатель влево
		} else {
			left++ // сумма слишком маленькая - двигаем левый указатель вправо
		}
	}

	return []int{left + 1, right + 1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
