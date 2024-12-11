package main

import (
	"fmt"
)

func main() {
	t := 12
	p := []int{10, 8, 0, 5, 3}
	s := []int{2, 4, 1, 1, 3}
	o := carFleet(t, p, s)
	fmt.Println(o)
}

func carFleet(target int, position []int, speed []int) int {
	fleets, travelTime := 0, 0.0
	road := make([]float64, target)

	for i, pos := range position {
		road[pos] = float64(target-pos) / float64(speed[i])
	}

	for i := target - 1; i >= 0; i-- {
		if road[i] > travelTime {
			travelTime = road[i]
			fleets++
		}
	}
	return fleets
}
