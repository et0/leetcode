package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	maxAverage := -10001.0
	size := len(nums)
	maxCurrent := 0.0

	for left, right := 0, 0; left < size && right < size; left++ {
		for ; right < size && right-left < k; right++ {
			// Добавляем элемент в общую сумму подмассива
			maxCurrent += float64(nums[right])
		}
		maxAverage = max(maxAverage, maxCurrent/float64(k))

		// Вычитаем "первый" элемент из подмассива, т.к. сдвигаем левую границу
		maxCurrent -= float64(nums[left])
	}

	return maxAverage
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 50))
}
