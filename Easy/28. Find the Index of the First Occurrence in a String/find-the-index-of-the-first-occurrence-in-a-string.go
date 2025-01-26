package main

import "fmt"

func strStr(haystack string, needle string) int {

	sizeNeedle, sizeHaystack := len(needle), len(haystack)

	for i := 0; i <= sizeHaystack-sizeNeedle; i++ {
		if haystack[i:i+sizeNeedle] == needle {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("c", "ccc"))
}
