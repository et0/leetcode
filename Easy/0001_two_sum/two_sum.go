package twosum

func TwoSumSlow(nums []int, target int) []int {
	len := len(nums)

	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func TwoSum(nums []int, target int) []int {
	maps := make(map[int]int)

	for k, v := range nums {
		check := target - v
		if exist, ok := maps[check]; ok {
			return []int{exist, k}
		}
		maps[v] = k
	}

	return []int{}
}
