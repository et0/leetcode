package le1550 // https://leetcode.com/problems/three-consecutive-odds/description/

func threeConsecutiveOdds(arr []int) bool {
	var (
		result bool
		three  int
	)

	for _, v := range arr {
		if v%2 == 1 {
			three++
			if three == 3 {
				result = true
				break
			}
			continue
		}

		three = 0
	}

	return result
}

func Wrapper(arr []int) bool {
	return threeConsecutiveOdds(arr)
}
