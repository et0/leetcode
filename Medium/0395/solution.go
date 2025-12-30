package main // https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/

func longestSubstring(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	// Кол-во уникальных символов
	uniq := make(map[byte]int, 26)
	for _, w := range s {
		uniq[byte(w)]++
	}

	// Стоп символы, число которых меньше K
	ignore := make(map[byte]struct{}, 26)
	for key, v := range uniq {
		if v < k {
			ignore[byte(key)] = struct{}{}
		}
	}

	// Если нет стоп символов, значит вся строка подходит
	if len(ignore) == 0 {
		return len(s)
	}

	maxLength := 0
	for left, right := 0, 0; left < len(s); {
		for ; right < len(s); right++ {
			if _, exist := ignore[s[right]]; exist {
				break
			}
		}

		if left == right {
			left++
			right = left
		} else if right-left < k {
			left = right
		} else {
			maxLength = max(maxLength, longestSubstring(s[left:right], k))
			left = right
		}
	}

	return maxLength
}
