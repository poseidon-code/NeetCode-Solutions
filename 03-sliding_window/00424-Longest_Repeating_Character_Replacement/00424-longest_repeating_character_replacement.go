// 424: Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func characterReplacement(s string, k int) int  {
	count := make([]int, 26)
	maxCount := 0
	i, j := 0, 0

	result := 0

	for j<len(s) {
		count[s[j]-'A']++;
		maxCount = int(math.Max(float64(maxCount), float64(count[s[j]-'A'])))

		if j-i+1-maxCount > k {
			count[s[i]-'A']--
			i++
		}
		result = int(math.Max(float64(result), float64(j-i+1)))
		j++
	}

	return result
}


func main() {
	// INPUT
	s := "ABAB"
	k := 2

	// OUTPUT
	result := characterReplacement(s, k)
	fmt.Println(result)
}