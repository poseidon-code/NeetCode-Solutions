// 20: Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/

package main

import "fmt"

// SOLUTION
func isValid(s string) bool {
	var parentheses []rune

	for _, c := range s {
		switch c {
			case '{':
				parentheses = append(parentheses, '}')
			case '[':
				parentheses = append(parentheses, ']')
			case '(':
				parentheses = append(parentheses, ')')
			default:
				if len(parentheses)==0 || c!=parentheses[len(parentheses)-1] {
					return false
				} else {
					parentheses = parentheses[:len(parentheses)-1]
				}
		}
	}

	return len(parentheses)==0
}


func main() {
	// INPUT
	s := "()[]{}"

	// OUTPUT
	result := isValid(s)
	fmt.Println(result)
}