package main

import "fmt"

func main() {
	t := Constructor()
	t.Set("foo", "bar", 1)
	fmt.Println(t.Get("foo", 1))
	fmt.Println(t.Get("foo", 3))
	t.Set("foo", "bar2", 4)
	fmt.Println(t.Get("foo", 4))
	fmt.Println(t.Get("foo", 5))
}

type TimeMap struct {
	KeyMap map[string][]Pair
}

type Pair struct {
	val       string
	timeStamp int
}

func Constructor() TimeMap {
	return TimeMap{
		KeyMap: make(map[string][]Pair),
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.KeyMap[key] = append(tm.KeyMap[key], Pair{value, timestamp})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	pairs, ok := tm.KeyMap[key]

	if !ok {
		return ""
	}

	left, right := 0, len(pairs)-1

	var val string
	for left <= right {
		mid := left + (right-left)/2
		if pairs[mid].timeStamp <= timestamp {
			val = pairs[mid].val
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return val
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
