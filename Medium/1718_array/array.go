package l_m_1718_array // https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/

func check(index int, n int, used map[int]int, seq []int) bool {
	for i := index; i < len(seq); i++ {
		if seq[i] != 0 {
			continue
		}
		for sizeN := n; sizeN > 1; sizeN-- {
			if _, ok := used[sizeN]; ok || i+sizeN > len(seq)-1 || seq[i+sizeN] != 0 {
				continue
			}

			used[sizeN] = sizeN
			seq[i], seq[i+sizeN] = sizeN, sizeN

			if len(used)*2+2 == len(seq)-1 || check(i+1, n, used, seq) {
				return true
			} else {
				delete(used, sizeN)
				seq[i], seq[i+sizeN] = 0, 0
			}
		}
	}

	return false
}

func constructDistancedSequence(n int) []int {
	seq := make([]int, n*2-1)

	// First element always N
	seq[0] = n
	if n > 1 {
		seq[n] = n
	}

	_ = check(1, n-1, make(map[int]int), seq)

	for k, v := range seq {
		if v == 0 {
			seq[k] = 1
			break
		}
	}

	return seq
}

func Wrapper(n int) []int {
	return constructDistancedSequence(n)
}
