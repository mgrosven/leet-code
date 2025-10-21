package main

import (
	"fmt"
)

func main() {
	rooms := [][]int{{1}, {2}, {3}, {}}
	result := canVisitAllRooms(rooms)
	fmt.Println(result) // Output: true

	rooms2 := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	result2 := canVisitAllRooms(rooms2)
	fmt.Println(result2) // Output: false
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))

	var visit func(room int)
	visit = func(room int) {
		if visited[room] {
			return
		}
		visited[room] = true
		for _, key := range rooms[room] {
			visit(key)
		}
	}

	visit(0)

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
