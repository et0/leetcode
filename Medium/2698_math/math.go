package l_m_2698_math

func checkSumSub(i, square, sum int) bool {
	j := 10
	for {
		div, mod := square/j, square%j
		// fmt.Println(i, square, sum, div, mod)

		if sum+div+mod == i {
			// fmt.Println("true1")
			return true
		}

		if mod >= square || div == 0 {
			// fmt.Println("false2")
			break
		}

		if checkSumSub(i, div, sum+mod) {
			// fmt.Println("true2")
			return true
		}

		j *= 10
	}

	return false
}

func punishmentNumber(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		pow := i * i
		if pow == i || checkSumSub(i, i*i, 0) {
			result += pow
		}
		// fmt.Println(i, pow, result, "\n")
	}

	return result
}

func Wrapper(n int) int {
	return punishmentNumber(n)
}
