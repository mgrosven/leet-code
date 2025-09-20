package main

import "fmt"

func main() {
	flowerbed := []int{1, 0, 0, 0}
	n := 1
	result := canPlaceFlowers(flowerbed, n)
	fmt.Println(result)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	free := 1

	for i, v := range flowerbed {
		if i == len(flowerbed)-1 {
			free++
		}
		if v == 0 {
			free++
			if free >= 3 {
				n--
				if n == 0 {
					return true
				}
				free = 1
			}
		} else {
			free = 0
		}
	}
	return false
}
