package l0026

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	unique := 0
	for i := 1; i < size; i++ {
		if nums[unique] == nums[i] {
			continue
		}
		unique++
		nums[unique] = nums[i]
	}

	return unique + 1
}

func Wrapper(nums []int) int {
	return removeDuplicates(nums)
}
