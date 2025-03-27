package le0191 // https://leetcode.com/problems/number-of-1-bits/description/

func hammingWeight(n int) int {
	bits := 1
	for n > 1 {
		mod := n / 2
		if mod*2 != n {
			bits++
		}
		n = mod
	}

	return bits
}

func Wrapper(n int) int {
	return hammingWeight(n)
}
