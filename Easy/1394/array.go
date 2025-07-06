package le1394 // https://leetcode.com/problems/find-lucky-integer-in-an-array/

func findLucky(arr []int) int {
	repeat := make(map[int]int)

	for _, a := range arr {
		repeat[a]++
	}

	max := -1
	for k, r := range repeat {
		if k == r && r > max {
			max = r
		}
	}

	return max
}

func Wrapper(arr []int) int {
	return findLucky(arr)
}
