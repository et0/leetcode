package le2379 // https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/description/

type option struct {
	size    int
	recolor int
}

func minimumRecolors(blocks string, k int) int {
	repeatB, min, startB, set := 0, 0, false, make([]option, 0, len(blocks))
	for _, v := range blocks {
		if v == 'B' {
			repeatB++
			if repeatB == k {
				return 0
			}

			if !startB {
				startB = true
				set = append(set, option{})
			}
		} else {
			startB = false
			repeatB = 0
			set = append(set, option{})
		}

		for i := 0; i < len(set); i++ {
			if set[i].size > k {
				continue
			}

			set[i].size++
			if v == 'W' {
				set[i].recolor++
			}

			if set[i].size == k {
				if min == 0 || set[i].recolor < min {
					min = set[i].recolor
				}
			}
		}
	}

	return min
}

func Wrapper(blocks string, k int) int {
	return minimumRecolors(blocks, k)
}
