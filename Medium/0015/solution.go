package main // https://leetcode.com/problems/3sum/

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}

	// Зачем сортировка?
	// Чтобы использовать два указателя и легко пропускать дубликаты.
	slices.Sort(nums)

	// Почему фиксируем только до len(nums)-2?
	// Потому что нужно минимум два элемента для left/right
	for fix := 0; fix < len(nums)-2; fix++ {
		// Если фикс. значение больше нуля, то дальше уже не получим сумму 0
		if nums[fix] > 0 {
			break
		}

		// Если фикс. значение равно предыдущей попытке
		if fix != 0 && nums[fix] == nums[fix-1] {
			continue
		}

		target := -nums[fix]
		for left, right := fix+1, len(nums)-1; left < right; {
			sum := nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[fix], nums[left], nums[right]})
				left++
				right--

				// Исключаем предыдущие дубли слева
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// Исключаем предыдущие дубли справа
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] > nums[fix]*-1 {
				right--
			} else {
				left++
			}
		}
	}

	return result
}

func main() {
	// [-4 1 2 2 3]
	threeSum([]int{-4, 1, 2, 2, 2, 2, 3})
}
