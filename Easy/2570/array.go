package le2570 // https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/description/

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	size1 := len(nums1)
	size2 := len(nums2)

	size := 0
	if size1 > size2 {
		size = size1
	} else {
		size = size2
	}
	nums := make([][]int, 0, size)

	i, j := 0, 0
	for {
		if i == size1 {
			if j < size2 {
				nums = append(nums, nums2[j:]...)
			}
			break
		} else if j == size2 {
			nums = append(nums, nums1[i:]...)
			break
		}

		if nums1[i][0] == nums2[j][0] {
			nums = append(nums, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i, j = i+1, j+1
		} else if nums1[i][0] > nums2[j][0] {
			nums = append(nums, nums2[j])
			j++
		} else {
			nums = append(nums, nums1[i])
			i++
		}
	}

	return nums
}

func Wrapper(nums1 [][]int, nums2 [][]int) [][]int {
	return mergeArrays(nums1, nums2)
}
