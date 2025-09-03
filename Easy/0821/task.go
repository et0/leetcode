package le0821 // https://leetcode.com/problems/shortest-distance-to-a-character/description/

// "loveleetcode", 'e'
func shortestToChar(s string, c byte) []int {
	answer := make([]int, len(s))

	counter := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			counter = 1
		} else if counter != 0 {
			answer[i] = counter
			counter++
		}
	}

	counter = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			counter = 1
		} else if counter != 0 {
			if answer[i] == 0 || answer[i] > counter {
				answer[i] = counter
			} else if answer[i] <= counter {
				continue
			}
			counter++
		}
	}

	return answer
}

func Wrapper(s string, c byte) []int {
	return shortestToChar(s, c)
}
