package le1504 // https://leetcode.com/problems/count-good-triplets/description/?

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	size := len(arr)
	good := 0
	for i := 0; i < size-2; i++ {
		for k := size - 1; k > i; k-- {
			if abs(arr[i]-arr[k]) > c {
				continue
			}

			for j := i + 1; j < k; j++ {
				if abs(arr[i]-arr[j]) > a {
					continue
				}
				if abs(arr[j]-arr[k]) > b {
					continue
				}
				good++
			}
		}
	}

	return good
}

func Wrapper(arr []int, a int, b int, c int) int {
	return countGoodTriplets(arr, a, b, c)
}
