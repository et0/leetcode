package main // https://leetcode.com/problems/move-zeroes/description/

import "fmt"

// Два указателя: left указывает на первый ноль, right сканирует массив.
// При нахождении ненулевого элемента меняем его с нулём на позиции left.
// Это сдвигает все нули вправо, сохраняя порядок ненулевых элементов.
func moveZeroes(nums []int) {
	left := 0

	for right := 0; right < len(nums); right++ {
		// Сдвигаем left вправо, пока он указывает на ненулевой элемент
		// (left всегда должен указывать на первый ноль в массиве)
		for ; left < right && nums[left] != 0; left++ {

		}

		// Если нашли ненулевой элемент - меняем его местами с первым нулём
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	fmt.Println(nums)
}

// Стратегия перезаписи: pos указывает, куда класть следующий ненулевой элемент.
// Первый проход: копируем все ненулевые элементы в начало.
// Второй проход: заполняем оставшиеся позиции нулями.
func moveZeroes2(nums []int) {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		// Копируем ненулевой элемент в начало массива
		nums[pos] = nums[i]
		pos++
	}

	// Заполняем оставшуюся часть массива нулями
	for ; pos < len(nums); pos++ {
		nums[pos] = 0
	}

	fmt.Println(nums)
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
