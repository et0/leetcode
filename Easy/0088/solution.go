package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n - 1; i >= 0; i-- {
		if n == 0 {
			nums1[i] = nums1[m-1]
			m--
		} else if m == 0 {
			nums1[i] = nums2[n-1]
			n--
		} else if nums1[m-1] >= nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}
	fmt.Println(nums1)
}

func main() {
	merge([]int{0, 0, 0}, 0, []int{2, 5, 6}, 3)
}
