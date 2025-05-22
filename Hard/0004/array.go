package lh0004 // https://leetcode.com/problems/median-of-two-sorted-arrays/description/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size1, size2 := len(nums1), len(nums2)
	size := size1 + size2
	half := size / 2
	median := struct {
		prev, current int
	}{}
	for i, j, step := 0, 0, 0; step <= half; step++ {
		if i == size1 {
			median.prev = median.current
			median.current = nums2[j]
			j++
		} else if j == size2 {
			median.prev = median.current
			median.current = nums1[i]
			i++
		} else {
			if nums1[i] < nums2[j] {
				median.prev = median.current
				median.current = nums1[i]
				i++
			} else {
				median.prev = median.current
				median.current = nums2[j]
				j++
			}
		}
	}

	if size%2 == 1 {
		return float64(median.current)
	}
	return float64(median.current+median.prev) / 2
}

func Wrapper(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}
