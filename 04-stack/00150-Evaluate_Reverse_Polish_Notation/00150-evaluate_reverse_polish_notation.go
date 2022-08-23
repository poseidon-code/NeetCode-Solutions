// 150: Evaluate Reverse Polish Notation
// https://leetcode.com/problems/evaluate-reverse-polish-notation/

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// SOLUTION
func evalRPN(tokens []string) int {
	var result []int

	for _, s := range tokens {
		if len(s)>1 || unicode.IsDigit(rune(s[0])) {
			si, _ := strconv.ParseInt(string(s), 10, 64)
			result = append(result, int(si))
		} else {
			x2 := result[len(result)-1]; result = result[:len(result)-1]
			x1 := result[len(result)-1]; result = result[:len(result)-1]
			
			switch s[0] {
				case '+': x1+=x2; break
				case '-': x1-=x2; break
				case '*': x1*=x2; break
				case '/': x1/=x2; break
			}

			result = append(result, x1)
		}
	}

	return result[len(result)-1]
}

func main() {
	// INPUT
	tokens := []string{"2","1","+","3","*"}

	// OUTPUT
	result := evalRPN(tokens)
	fmt.Println(result)
}