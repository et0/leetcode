package le3375 // https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k/description/

func minOperations(nums []int, k int) int {
	db := make(map[int]bool, len(nums))

	for _, v := range nums {
		if v < k {
			return -1
		}

		db[v] = true
	}

	db[k] = true

	return len(db) - 1
}

func Wrapper(nums []int, k int) int {
	return minOperations(nums, k)
}
