package main

import (
	"fmt"
	"math"
)

func main() {
	asteroids := []int{5, 10, -5}
	result := asteroidCollision(asteroids)
	fmt.Println(result) // Output: [5, 10]

	asteroids = []int{8, -8}
	result = asteroidCollision(asteroids)
	fmt.Println(result) // Output: []

	asteroids = []int{10, 2, -5}
	result = asteroidCollision(asteroids)
	fmt.Println(result) // Output: [10]

	asteroids = []int{-2, -1, 1, 2}
	result = asteroidCollision(asteroids)
	fmt.Println(result) // Output: [-2, -1, 1, 2]

	asteroids = []int{3, 5, -6, 2, -1, 4}
	result = asteroidCollision(asteroids)
	fmt.Println(result) // Output: [-2, -2]
}

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, asteroid := range asteroids {
		if len(stack) == 0 || asteroid > 0 {
			stack = append(stack, asteroid)
			continue
		}

		for {
			top := stack[len(stack)-1]
			if top < 0 {
				stack = append(stack, asteroid)
				break
			}

			if math.Abs(float64(top)) == math.Abs(float64(asteroid)) {
				stack = stack[:len(stack)-1]
				break
			} else if math.Abs(float64(top)) > math.Abs(float64(asteroid)) {
				break
			} else {
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					stack = append(stack, asteroid)
					break
				}
			}
		}
	}

	return stack
}
