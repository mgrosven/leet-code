package main

import (
	"fmt"
)

func main() {

	n := 4
	result := tribonacci(n)
	fmt.Println(result) // Output: 4

	n = 25
	result = tribonacci(n)
	fmt.Println(result) // Output: 1389537
}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	a, b, c := 0, 1, 1
	var d int
	for i := 3; i <= n; i++ {
		d = a + b + c
		a, b, c = b, c, d
	}
	return d
}
