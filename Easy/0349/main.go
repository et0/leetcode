package le0349 // https://leetcode.com/problems/intersection-of-two-arrays/description/

import (
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	size1 := len(nums1)
	size2 := len(nums2)

	data := make(map[int]struct{}, min(size1, size2))
	for _, v := range nums1 {
		if !search(v, &nums2) {
			continue
		}

		if _, ok := data[v]; ok {
			continue
		}

		data[v] = struct{}{}
	}

	out := make([]int, 0, len(data))
	for i := range data {
		out = append(out, i)
	}

	return out
}

func search(n int, nums *[]int) bool {
	left, right := 0, len(*nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if (*nums)[mid] == n {
			return true
		}

		if (*nums)[mid] > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

func Wrapper(nums1 []int, nums2 []int) []int {
	return intersection(nums1, nums2)
}
