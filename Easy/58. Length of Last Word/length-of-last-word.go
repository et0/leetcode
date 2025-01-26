package main

import "fmt"

func lengthOfLastWord(s string) int {
	size := len(s)
	index := 0
	for i := size - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if index != 0 {
				return index
			}
			continue
		}
		index++
	}

	return index
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
