package main // https://leetcode.com/problems/container-with-most-water/

import "fmt"

// maxArea находит максимальный объём воды, который можно вместить
// между двумя линиями массива height. Использует метод двух указателей:
// на каждом шаге сдвигаем указатель с меньшей высотой внутрь.
// O(n) времени, O(1) памяти.
func maxArea(height []int) int {
	// Указатели на края массива: максимально возможная ширина
	left, right := 0, len(height)-1

	// Текущий максимум объёма
	maxAmount := 0

	for left < right {
		// Считаем объём для текущей пары:
		// высота ограничена меньшим столбиком, ширина — расстоянием между ними
		width := right - left
		h := min(height[left], height[right])
		maxAmount = max(maxAmount, width*h)

		// Сдвигаем меньший столбик внутрь:
		// только это может увеличить высоту контейнера.
		// Если сдвинуть больший — высота не изменится, а ширина уменьшится.
		if height[left] < height[right] {
			left++
		} else if height[left] > height[right] {
			right--
		} else {
			// Равные столбики — оба ограничивают высоту,
			// можно сдвинуть оба без потери оптимального решения
			left++
			right--
		}
	}

	return maxAmount
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
