package le2843 // https://leetcode.com/problems/count-symmetric-integers/

func countSymmetricIntegers(low int, high int) int {
	symmetric := 0

	if low < 10 {
		low = 10
	}
	for ; low <= high && low < 100; low++ {
		if low/10 == low%10 {
			symmetric++
		}
	}

	if low > 99 && low < 1000 {
		low = 1000
	}

	for ; low <= high && low != 10000; low++ {
		if low%10+low%100/10 == low%1000/100+low%10000/1000 {
			symmetric++
		}
	}

	return symmetric
}

func Wrapper(low int, high int) int {
	return countSymmetricIntegers(low, high)
}
