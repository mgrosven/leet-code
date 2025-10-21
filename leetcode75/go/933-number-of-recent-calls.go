package main

import (
	"fmt"
)

func main() {
	rc := Constructor()
	fmt.Println(rc.Ping(1))    // return 1
	fmt.Println(rc.Ping(100))  // return 2
	fmt.Println(rc.Ping(3001)) // return 3
	fmt.Println(rc.Ping(3002)) // return 3
}

type RecentCounter struct {
	ping []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		ping: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.ping = append(this.ping, t)
	lowerBound := func(arr []int, target int) int {
		for i, v := range arr {
			if v >= target {
				return i
			}
		}
		return len(arr)
	}
	return len(this.ping) - lowerBound(this.ping, t-3000)
}
