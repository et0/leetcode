package array // https://leetcode.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {
	size := len(s)
	result := make([]byte, 0, size)
	step := numRows*2 - 2
	if step == 0 {
		step = 1
	}
	for i := 0; i < numRows; i++ {
		for j := i; j < size; j += step {
			result = append(result, s[j])

			if i != 0 && i+1 != numRows && j+step-i*2 < size {
				result = append(result, s[j+step-i*2])
			}
		}
	}

	return string(result[:])
}

func Wrapper(s string, numRows int) string {
	return convert(s, numRows)
}
