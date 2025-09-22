package le3005 // https://leetcode.com/problems/count-elements-with-maximum-frequency/

func maxFrequencyElements(nums []int) int {
	freq := make(map[int]int)
	maxFreq := 1
	out := 0

	for _, v := range nums {
		freq[v]++

		count := freq[v]
		if count == maxFreq {
			out += count
		} else if count > maxFreq {
			maxFreq = count
			out = count
		}
	}

	return out
}
