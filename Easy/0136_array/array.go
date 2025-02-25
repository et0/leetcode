package array // https://leetcode.com/problems/single-number/

func singleNumberMap(nums []int) int {
	numbers := make(map[int]bool, len(nums)/2+1)

	for _, v := range nums {
		if _, ok := numbers[v]; ok {
			delete(numbers, v)
		} else {
			numbers[v] = true
		}
	}

	for k, _ := range numbers {
		return k
	}

	return 0
}

func singleNumber(nums []int) int {
	result := 0

	for _, v := range nums {
		result ^= v
	}

	return result
}

func Wrapper(nums []int) int {
	return singleNumber(nums)
}
