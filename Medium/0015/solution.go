package main // https://leetcode.com/problems/3sum/

import (
	"slices"
)

// Идея: сортируем; фиксируем первое число one, а пару к нему ищем как в #167 — two с начала
// остатка, three с конца, target равен -nums[one]. Если nums[one] больше нуля, дальше все
// числа тоже положительные и ноль не набрать — выходим.
// Дубли: one пропускает значения, равные предыдущему; после найденной тройки сдвигаем оба
// указателя (при том же one новое значение two требует другого партнёра, со старым three
// в ноль не сложится) и пропускаем повторы two и three.
// Время: O(n^2) — до n значений one, для каждого внутренний проход O(n): зазор three-two
// сужается на 1 за шаг, пропуски дублей тоже сдвигают указатели. Сортировка O(n log n) теряется.
// Память: O(log n) дополнительно — стек pdqsort в slices.Sort, сортировка на месте; сам
// алгоритм O(1). Ответ не считаем: троек бывает до O(n^2).
// Ловушка: ветка "нашли" должна быть частью одной цепочки if / else if / else и сдвигать
// оба указателя — иначе бесконечный цикл или повторы; условие для one — меньше или равно
// нулю, иначе теряется [0,0,0].
func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	out := [][]int{}

	size := len(nums)
	for one := 0; one < size && nums[one] <= 0; one++ {
		if one > 0 && nums[one] == nums[one-1] {
			continue
		}

		for two, three := one+1, size-1; two < three; {
			if nums[one]+nums[two]+nums[three] == 0 {
				out = append(out, []int{nums[one], nums[two], nums[three]})

				two++
				three--

				for two < three && nums[two] == nums[two-1] {
					two++
				}

				for two < three && nums[three] == nums[three+1] {
					three--
				}
			} else if nums[two]+nums[three] > -nums[one] {
				three--
			} else {
				two++
			}
		}
	}

	return out
}
