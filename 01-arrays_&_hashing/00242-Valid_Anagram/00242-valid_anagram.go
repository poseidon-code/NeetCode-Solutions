// 242: Valid Anagram
// https://leetcode.com/problems/valid-anagram/

package main

import "fmt"

func isAnagram(s, t string) bool {
    if len(s) != len(t) { return false }

    n := len(s)
    counts := [26]int{}

    for i:=0; i<n; i++ {
        counts[s[i] - 'a']++
        counts[t[i] - 'a']--
    }

    for i:=0; i<26; i++ {
        if counts[i] != 0 {return false}
    }

    return true
}

func main() {
	// INPUT
	s := "rat"
    t := "cat"

	// OUTPUT
    result := isAnagram(s, t)
    fmt.Println(result)
}
