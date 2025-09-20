package main

import "fmt"

func main() {

	s := "hello"
	result := reverseVowels(s)
	fmt.Println(result)
}

func reverseVowels(s string) string {

	bytes := []byte(s)
	left, right := 0, len(bytes)-1

	for left < right {
		for left < right && !isVowel(bytes[left]) {
			left++
		}
		for left < right && !isVowel(bytes[right]) {
			right--
		}
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
	return string(bytes)

}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
