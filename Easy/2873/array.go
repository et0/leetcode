package le2873 // https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-i/description/

func maximumTripletValue(nums []int) int64 {
	var maxElement, maxDiff, maxTriplet int64

	size := len(nums)
	for i := 0; i < size; i++ {
		maxTriplet = max(maxTriplet, maxDiff*int64(nums[i]))
		maxElement = max(maxElement, int64(nums[i]))
		maxDiff = max(maxDiff, maxElement-int64(nums[i]))
	}
	if maxTriplet > 0 {
		return maxTriplet
	}
	return 0
}

func Wrapper(nums []int) int64 {
	return maximumTripletValue(nums)
}
