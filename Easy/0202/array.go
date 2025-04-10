package le0202 // https://leetcode.com/problems/happy-number/description/

func isHappy(n int) bool {
	squares := map[int]int{0: 0, 1: 1, 2: 4, 3: 9, 4: 16, 5: 25, 6: 36, 7: 49, 8: 64, 9: 81}

	for n > 9 {
		tmp := 0
		for n != 0 {
			newN := n / 10
			tmp += squares[n-newN*10]
			n = newN
		}
		n = tmp
	}

	return n == 1 || n == 7
}

func Wrapper(n int) bool {
	return isHappy(n)
}
