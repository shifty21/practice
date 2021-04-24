/**
 * @input A : Integer
 *
 * @Output string array.
 */
//  https://www.interviewbit.com/problems/generate-all-parentheses-ii/#
package main

func generateParenthesis(A int) []string {
	var result []string
	output := ""
	open := A
	close := A
	return solve(open, close, output, result)
}

func solve(open, close int, output string, result []string) []string {
	if open == 0 && close == 0 {
		result = append(result, output)
		return result
	}

	if open != 0 {
		op1 := output + "("
		result = solve(open-1, close, op1, result)
	}
	if close > open {
		op2 := output + ")"
		result = solve(open, close-1, op2, result)
	}
	return result
}
