// 22: Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/

package main

import "fmt"

// SOLUTION
func generateParentheses(n int) []string {
	var result []string
	generate(n, 0, 0, "", &result)
	return result
}

func generate(n, open, close int, s string, result *[]string) {
	if open==n && close==n {
		*result = append(*result, s)
		return
	}
	if (open < n) {
		generate(n, open+1, close, s+"(", result);
	}
	if (open > close) {
		generate(n, open, close+1, s+")", result);
	}
}


func main() {
	// INPUT
	n := 3

	// OUTPUT
	result := generateParentheses(n)
	fmt.Println(result)
}