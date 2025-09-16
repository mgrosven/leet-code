package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	r := maxProfit(a)
	fmt.Println(r)
}

func maxProfit(prices []int) int {
	minP, profit := prices[0], 0

	for _, p := range prices {
		minP = min(p, minP)
		profit = max(profit, p-minP)
	}
	return profit
}
