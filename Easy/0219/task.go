package le2019 // https://leetcode.com/problems/contains-duplicate-ii/description/

func containsNearbyDuplicate(nums []int, k int) bool {
	uniq := make(map[int]int)

	for i, v := range nums {
		value, ok := uniq[v]
		if !ok {
			uniq[v] = i

			continue
		}

		if i-value <= k {
			return true
		} else {
			uniq[v] = i
		}
	}

	return false
}

func Wrapper(nums []int, k int) bool {
	return containsNearbyDuplicate(nums, k)
}
