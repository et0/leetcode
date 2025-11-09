package le2169 // https://leetcode.com/problems/count-operations-to-obtain-zero/

func countOperations(num1 int, num2 int) int {
	counter := 0
	for num1 != 0 && num2 != 0 {
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}

		counter++
	}

	return counter
}

func Wrapper(num1 int, num2 int) int {
	return countOperations(num1, num2)
}
