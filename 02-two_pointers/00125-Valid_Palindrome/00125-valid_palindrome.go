// 125: Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isalnum(c string) bool {
    return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(c)
}

// SOLUTION
func isPalindrome(s string) bool {
	l := 0;
	r := len(s) - 1

	for l<r {
		for (!isalnum(string(s[l])) && l<r) {l++}
		for (!isalnum(string(s[r])) && l<r) {r--}
		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {return false}

		l++
		r--
	}

	return true
}


func main() {
	// INPUT
	s := "race a car"

	// OUTPUT
	result := isPalindrome(s)
	fmt.Println(result)
}