package main

import (
	"fmt"
	"log"
)

func Print(val int) {
	if val == 1 {
		log.Println(val)
		return
	}
	log.Println(val)
	Print(val - 1)
}

func main() {
	Print(7)
}

func removeOuterParentheses(S string) string {
	var result []string
	left := '('
	right := ')'
	left_count := 0
	right_count := 0
	temp := ""
	for i, char := range S {
		fmt.Printf("char at index %v value=%v", i, char)
		if char == left {
			left_count = left_count + 1
		}
		if char == right {
			right_count = right_count + 1
		}
		temp = temp + string(char)
		if left_count == right_count {
			result = append(result, string(char))
		}

	}
	fmt.Printf("%v", result)
	return ""
}
