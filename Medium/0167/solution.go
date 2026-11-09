package main // https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

// Идея: указатели с двух концов; сумма больше target — сдвигаем right влево (сумма уменьшится),
// меньше — сдвигаем left вправо (сумма увеличится). У каждого хода ровно один эффект на сумму.
// Время: O(n) — каждый шаг сужает зазор right-left на 1, шагов не больше n-1.
// Память: O(1) — только два указателя, результат из 2 элементов от n не зависит.
// Ловушка: оба указателя слева и только вправо (окно) — когда right растёт, нужный партнёр
// уменьшается, а left назад не умеет; ломается на [1,2,3,4,4,9], target 8.
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			break
		}
		if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{left + 1, right + 1}
}
