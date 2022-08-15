// 3: Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func lengthOfLongestSubstring(s string) int {
	letters := make(map[byte]int)
	i, j := 0, 0
	result := 0

	for i<len(s) {
		if _, ok := letters[s[i]]; ok {
			j = int(math.Max(float64(j), float64(letters[s[i]]+1)))
		}
		letters[s[i]] = i
		result = int(math.Max(float64(result), float64(i-j+1)))
		i++
	}

	return result
}

func main() {
	// INPUT
	s := "abcabcbb"

	// OUTPUT
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}