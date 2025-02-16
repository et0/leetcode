// https://leetcode.com/problems/tuple-with-same-product/
package l_m_1726_array

import (
	"fmt"
	"sort"
)

func tupleSameProduct(nums []int) int {
	result := 0

	size := len(nums)
	if size < 4 {
		return result
	}

	ab := make(map[int]int, size)
	for ia := 0; ia < size; ia++ {
		for ib := ia + 1; ib < size; ib++ {
			ab[nums[ia]*nums[ib]]++
		}
	}
	fmt.Println(ab)

	for _, v := range ab {
		result += (v - 1) * v * 4
	}

	return result
}

func tupleSameProductBrute(nums []int) int {
	size := len(nums)
	if size < 4 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	hash := make(map[int]int, size)
	for k, v := range nums {
		hash[v] = k
	}

	result := 0
	for ia := 0; ia < size; ia++ {
		//fmt.Println("a ", nums[ia])
		for ib := ia + 1; ib < size; ib++ {
			//fmt.Println("b ", nums[ib], " ")
			ab := nums[ia] * nums[ib]

			repeat := make(map[int]int, size)
			for ic := ib + 1; ic < size; ic++ {

				// fmt.Println("c ", nums[ic], " ")
				if ic == ia || ic == ib || ab < nums[ic] || ab%nums[ic] != 0 || ab/nums[ic] == nums[ic] {
					continue
				}
				if _, ok := repeat[nums[ic]]; ok {
					continue
				}

				if id, ok := hash[ab/nums[ic]]; ok {
					// fmt.Println(nums[ia], nums[ib], nums[ic], ab/nums[ic])
					result += 8
					repeat[ab/nums[ic]] = id
					// break
				}
			}

		}
		//fmt.Println("")
	}

	//fmt.Println(hash)

	return result
}

func Wrapper(nums []int) int {
	return tupleSameProduct(nums)
}
