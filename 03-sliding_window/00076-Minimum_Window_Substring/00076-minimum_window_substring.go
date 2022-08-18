// 76: Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/

package main

import "fmt"

// SOLUTION
func minWindow(s, t string) string {
	m := make(map[byte]int)
	for i:=0; i<len(t); i++ {m[t[i]]++}

	i := 0
	j := 0
	counter := len(t)
	
	minStart := 0
	minLength := len(s)+1
	
	for j < len(s) {
		if m[s[j]] > 0 {counter--}
		m[s[j]]--
		j++

		for counter==0 {
			if j-i < minLength {
				minStart = i
				minLength = j-i
			}
			
			m[s[i]]++
			if m[s[i]] > 0 {counter++}
			i++
		}
	}
	
	if (minLength != len(s)) {
		return s[minStart:minStart+minLength]
	}
	return ""
}


func main() {
	// INPUT
	s := "ADOBECODEBANC"
    t := "ABC"

	// OUTPUT
	result := minWindow(s, t)
	fmt.Println(result)
}