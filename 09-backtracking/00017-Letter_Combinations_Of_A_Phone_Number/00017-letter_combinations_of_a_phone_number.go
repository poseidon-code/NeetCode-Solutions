// 17: Letter Combinations Of A Phone Number
// https://leetcode.com/problems/letter-combinations-of-a-phone-number

package main

import "fmt"

func backtrack(digits string, index int, m map[byte]string, current string, result *[]string) {
	if index == len(digits) {
		*result = append(*result, current)
		return
	}

	s := m[digits[index]]
	for i := 0; i < len(s); i++ {
		current += string(s[i])
		backtrack(digits, index+1, m, current, result)
		current = current[:len(current)-1]
	}
}

// SOLUTION
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var current string
	var result []string
	backtrack(digits, 0, m, current, &result)
	return result
}

func main() {
	// INPUT
	digits := "23"

	// OUTPUT
	result := letterCombinations(digits)

	fmt.Println(result)
}
