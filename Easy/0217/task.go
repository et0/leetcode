package le0217 // https://leetcode.com/problems/contains-duplicate/description/

func containsDuplicate(nums []int) bool {
	uniq := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := uniq[v]; ok {
			return true
		} else {
			uniq[v] = struct{}{}
		}
	}

	return false
}

func Wrapper(nums []int) bool {
	return containsDuplicate(nums)
}
