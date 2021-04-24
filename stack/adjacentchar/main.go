package main

import (
	"fmt"
)

func removeDuplicates(S string) string {
	var data []rune
	for _, c := range S {
		if len(data) == 0 {
			data = append(data, c)
			continue
		}
		if c == data[len(data)-1] {
			data = data[:len(data)-1]
			continue
		}
		data = append(data, c)
	}
	return string(data)
}

func main() {
	fmt.Printf("a=%d\n", int('a'))
}

func makeGood(s string) string {
	var result []rune
	var diff int
	result = append(result, rune(s[0]))
	for i := 1; i < len(s)-1; i++ {
		if rune(s[i]) > result[len(result)-1] {
			diff = int(rune(s[i]) - result[len(result)-1])
		} else {
			diff = int(result[len(result)-1] - rune(s[i]))
		}
		if diff == 32 {
			i = i - 1
			result = result[:len(result)-1]
			continue
		}
		result = append(result, rune(s[i]))

	}
	fmt.Printf("Result=%v", string(result))
	return string(result)
}
