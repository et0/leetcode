package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	remove := regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	s = strings.ToLower(remove.ReplaceAllString(s, ""))

	chars := []rune{}
	for _, v := range s {
		chars = append(chars, v)
	}

	length := len(chars)
	for i := 0; i < length/2; i++ {
		if chars[i] != chars[length-i-1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("мамвам"))
}
