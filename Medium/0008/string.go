package lm0008_string // https://leetcode.com/problems/string-to-integer-atoi/description/

import (
	"math"
)

func myAtoi(s string) int {
	var (
		started, digital, minused, nulled bool
	)
	result := make([]int, 0, 200)

	for _, v := range s {
		if v == ' ' {
			if started {
				break
			}
		} else if v == '-' {
			if started || nulled {
				break
			}
			started, minused = true, true
		} else if v == '+' {
			if started || nulled {
				break
			}
			started = true
		} else if v == '0' {
			if !started {
				started = true

				if nulled {
					continue
				} else {
					nulled = true
				}
			} else if digital {
				result = append(result, 0)
			}
		} else if v == '1' || v == '2' || v == '3' || v == '4' || v == '5' || v == '6' || v == '7' || v == '8' || v == '9' {
			switch v {
			case '1':
				result = append(result, 1)
			case '2':
				result = append(result, 2)
			case '3':
				result = append(result, 3)
			case '4':
				result = append(result, 4)
			case '5':
				result = append(result, 5)
			case '6':
				result = append(result, 6)
			case '7':
				result = append(result, 7)
			case '8':
				result = append(result, 8)
			case '9':
				result = append(result, 9)
			}

			if !started {
				started = true
			}
			if !digital {
				digital = true
			}
		} else {
			break
		}
	}

	var resultInt int
	size := len(result) - 1
	if size > 10 {
		if minused {
			return -2147483648
		}
		return 2147483647
	}

	for k, v := range result {
		if v == 0 {
			continue
		}
		resultInt += int(math.Pow10(size-k)) * v
	}

	if minused {
		resultInt *= -1
	}

	if resultInt < -2147483648 {
		return -2147483648
	} else if resultInt > 2147483647 {
		return 2147483647
	}

	return resultInt
}

func Wrapper(s string) int {
	return myAtoi(s)
}
