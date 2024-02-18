// 131: Palindrome Partitioning
// https://leetcode.com/problems/palindrome-partitioning

package main

import "fmt"


func backtrack(s string, start int, current []string, result *[][]string) {
	if start == len(s) {
		t := make([]string, len(current))
		copy(t, current)
		*result = append(*result, t)
		return
	}
	
	for i:=start; i<len(s); i++ {
		if (isPalindrome(s, start, i)) {
			str := s[start:i+1]
			current = append(current, str)
			backtrack(s, i + 1, current, result)
			current = current[:len(current)-1]
		}
	}
}

func isPalindrome(s string, start, end int) bool {
    for start <= end {
        if s[start] != s[end] {
            return false
        }
        start++
        end--
    }
    return true
}


// SOLUTION
func partition(s string) [][]string {
    var current []string
	var result [][]string
	backtrack(s, 0, current, &result)
	return result;
}

func main() {
    // INPUT
    s := "aab"

    // OUTPUT
    result := partition(s)

	fmt.Println(result)
}
