package main

import (
	"fmt"
)

func main() {
	n := 1
	o := generateParenthesis(n)
	fmt.Println(o)
}

func generateParenthesis(n int) []string {
	result := []string{}

	var generate func(s string, o, c int)
	generate = func(s string, o, c int) {
		if len(s) == 2*n {
			result = append(result, s)
			return
		}

		if o < n {
			generate(s+"(", o+1, c)
		}
		if c < o {
			generate(s+")", o, c+1)
		}
	}

	generate("", 0, 0)
	return result
}
