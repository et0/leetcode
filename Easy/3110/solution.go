package main // https://leetcode.com/problems/score-of-a-string/description/

func scoreOfString(s string) int {
	score := 0
	length := len(s)

	for i := 0; i < length-1; i++ {
		sub := int(s[i]) - int(s[i+1])
		if sub < 0 {
			sub *= -1
		}
		score += sub
	}

	return score
}
