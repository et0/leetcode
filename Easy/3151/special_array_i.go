package specialarrayi

func isArraySpecial(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%2 == 0 && nums[i+1]%2 == 1 {
			continue
		} else if nums[i]%2 != 0 && nums[i+1]%2 != 1 {
			continue
		}
		return false
	}

	return true
}

func Wrapper(nums []int) bool {
	return isArraySpecial(nums)
}
