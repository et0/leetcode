// https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/

package l_e_1790_string

func areAlmostEqual(s1 string, s2 string) bool {
	flag := make([]int, 0, 2)

	for k, v := range s1 {
		if v != rune(s2[k]) {
			if len(flag) == 2 {
				return false
			}

			flag = append(flag, k)
			if len(flag) == 1 {
				continue
			} else {
				if s1[flag[0]] == s2[flag[1]] && s1[flag[1]] == s2[flag[0]] {
					continue
				} else {
					return false
				}
			}
		}
	}

	return len(flag) == 0 || len(flag) == 2
}

func Wrapper(s1 string, s2 string) bool {
	return areAlmostEqual(s1, s2)
}
