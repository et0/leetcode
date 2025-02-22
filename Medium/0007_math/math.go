package math

import (
	"math"
)

func reverse(x int) int {
	var res int

	flag := false
	if x < 0 {
		x *= -1
		flag = true
	}

	slice := make([]int, 0, 10)
	for i := 1; i <= x; i *= 10 {
		number := (x % (i * 10)) / i
		if number == 0 && len(slice) == 0 {
			continue
		}
		slice = append(slice, number)
	}

	size := float64(len(slice) - 1)
	for _, v := range slice {
		res += v * int(math.Pow(10.0, size))
		size -= 1.0
	}

	if res > 2147483647 || res < -2147483648 {
		return 0
	}

	if flag {
		res *= -1
	}

	return res
}

func Wrapper(x int) int {
	return reverse(x)
}
