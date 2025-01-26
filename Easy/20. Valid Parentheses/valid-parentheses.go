package main

import "fmt"

func isValid(s string) bool {
	close := map[string]string{")": "(", "}": "{", "]": "["}
	open := map[string]string{"(": ")", "{": "}", "[": "]"}

	stack := make([]string, 0, len(s))

	for _, v := range s {
		// скобка открывается?
		if _, ok := open[string(v)]; ok {
			stack = append(stack, string(v))
		} else if len(stack) > 0 && close[string(v)] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()({[]})[]"))
}
