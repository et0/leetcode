package main // https://leetcode.com/problems/range-sum-query-immutable/

// Идея: один раз считаем префиксные суммы: prefix[i] — сумма первых i элементов,
// prefix[0] = 0. Тогда сумма на отрезке [left, right] — это prefix[right+1] - prefix[left].
// Предподсчёт окупается на множестве запросов: O(n + q) вместо O(n * q), если считать
// каждую сумму циклом (q — число вызовов SumRange).
// Время: Constructor — O(n), один проход по массиву; SumRange — O(1), два обращения по индексу.
// Память: O(n) — массив prefix из n+1 элементов; SumRange дополнительной памяти не выделяет.
// Ловушка: в цикле складывать с предыдущим префиксом, а не с предыдущим элементом nums.
// Массив длины n+1 с нулём в начале убирает особый случай left == 0.
type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	data := NumArray{
		prefix: make([]int, len(nums)+1),
	}

	for i, v := range nums {
		data.prefix[i+1] = v + data.prefix[i]
	}

	return data
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix[right+1] - this.prefix[left]
}
