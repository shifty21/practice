package main

import (
	"strings"
	"unicode"
)

func letterCasePermutation(S string) []string {
	result := []string{}
	if len(S) == 0 {
		result = append(result, S)
		return result
	}
	output := ""
	solve(&result, output, S)
	return result
}

func solve(result *[]string, output, S string) {
	if len(S) == 0 {
		*result = append(*result, output)
		return
	}
	var op1, op2 string
	if unicode.IsNumber(rune(S[0])) {
		op1 = output + string(S[0])
		solve(result, op1, S[1:])
		return
	} else {
		op1 = output + strings.ToLower(string(S[0]))
		op2 = output + strings.ToUpper(string(S[0]))
	}
	S = S[1:]
	solve(result, op1, S)
	solve(result, op2, S)

}
