package le0169 // https://leetcode.com/problems/majority-element/description/

func majorityElement(nums []int) int {
	majority := make(map[int]int)
	half := len(nums) / 2
	for _, v := range nums {
		majority[v]++
		if majority[v] > half {
			return v
		}
	}

	return 0
}

func Wrapper(nums []int) int {
	return majorityElement(nums)
}
