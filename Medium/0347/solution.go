package main // https://leetcode.com/problems/top-k-frequent-elements/

// Идея: считаем, сколько раз встречается каждое число, и раскладываем числа по
// корзинам, где индекс корзины — это частота. Дальше идём по корзинам с конца,
// от самой большой частоты, и набираем k чисел. Сортировка не нужна: частота
// ограничена длиной массива, поэтому корзины сами дают нужный порядок.
// Время: O(n) — проход по nums, проход по map (в ней не больше n ключей) и проход
// по корзинам (их n+1); внутренние итерации в сумме дают не больше k.
// Требование условия "быстрее, чем O(n log n)" выполнено.
// Память: O(n) — в map не больше n ключей, в корзинах суммарно столько же чисел.
// Ловушка: индекс корзины — это частота, а не значение числа. Поэтому размер массива
// корзин берётся от длины входа, len(nums)+1. Если взять его из диапазона значений
// (20001), то на входе, где одно число встречается чаще 20000 раз, будет паника.
func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int, len(nums))

	for _, num := range nums {
		counts[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, val := range counts {
		buckets[val] = append(buckets[val], num)
	}

	out := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		if len(buckets[i]) == 0 {
			continue
		}

		for _, b := range buckets[i] {
			out = append(out, b)
			if len(out) >= k {
				return out
			}
		}

	}

	return out
}
