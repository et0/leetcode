package le3024 // https://leetcode.com/problems/type-of-triangle/description/

func triangleType(nums []int) string {
	if nums[0]+nums[1] <= nums[2] || nums[0]+nums[2] <= nums[1] || nums[1]+nums[2] <= nums[0] {
		return "none"
	}

	if nums[0] == nums[1] && nums[0] == nums[2] && nums[1] == nums[2] {
		return "equilateral"
	} else if nums[0] == nums[1] || nums[0] == nums[2] || nums[1] == nums[2] {
		return "isosceles"
	} else if nums[0]+nums[1] > nums[2] && nums[0]+nums[2] > nums[1] && nums[1]+nums[2] > nums[0] {
		return "scalene"
	}

	return "none"
}

func Wrapper(nums []int) string {
	return triangleType(nums)
}
