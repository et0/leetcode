package lm2308 // https://leetcode.com/problems/alternating-groups-ii/description/

func numberOfAlternatingGroups(colors []int, k int) int {
	var (
		start, end int
		firstEnd   int

		size   int = len(colors)
		groups int
	)

	for i := 1; i < size; i++ {
		if colors[i] == colors[i-1] {
			if i-start >= k {
				groups += end - start - k + 2
			}
			if start == 0 {
				firstEnd = i - 1
			}
			start = i
		} else {
			end = i
		}
	}

	if colors[size-1] != colors[0] {
		if start == 0 {
			groups += size
		} else {
			for start < size {
				if size-start >= k {
					groups++
				} else if k-(size-start) <= firstEnd+1 {
					groups++
				} else {
					break
				}
				start++
			}
		}
	} else {
		if delta := end - start - k + 2; delta > 0 {
			groups += delta
		}
	}

	return groups
}

func Wrapper(colors []int, k int) int {
	return numberOfAlternatingGroups(colors, k)
}
