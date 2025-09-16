package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	r := evalRPN(tokens)
	fmt.Println(r)
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, t := range tokens {
		v, err := strconv.Atoi(t)
		if err == nil {
			stack = append(stack, v)
		} else {
			// if stack.Size() < 2 ?? Not a valid senario?
			v1 := stack[len(stack)-1]
			v2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch t {
			case "+":
				stack = append(stack, v2+v1)
			case "-":
				stack = append(stack, v2-v1)
			case "*":
				stack = append(stack, v2*v1)
			case "/":
				stack = append(stack, v2/v1)
			}
		}
	}
	return stack[0]
}
