package main

import (
	"fmt"
)

func main() {

	taskList := []byte{'A', 'C', 'A', 'B', 'D', 'B'}
	n := 1
	fmt.Println(leastInterval(taskList, n))
}

// (max - 1) * (n + 1) + numMaxCount
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	cnt := make([]int, 26)
	for _, task := range tasks {
		cnt[task-'A']++
	}

	max, numMax := 0, 0
	for _, c := range cnt {
		if c > max {
			max = c
			numMax = 1
		} else if c == max {
			numMax++
		}
	}
	res := (n+1)*(max-1) + numMax
	if res > len(tasks) {
		return res
	} else {
		return len(tasks)
	}
}
