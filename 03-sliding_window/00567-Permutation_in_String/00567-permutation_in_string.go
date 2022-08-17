// 567: Permutation in String
// https://leetcode.com/problems/permutation-in-string/

package main

import (
	"fmt"
	"reflect"
)

// SOLUTION
func checkInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false;
	}
        
	s1Count := make([]int, 26); for _, c := range s1 { s1Count[c-'a']++ }
	s2Count := make([]int, 26)
	i, j := 0, 0
	
	
	for j < len(s2) {
		s2Count[s2[j]-'a']++
		
		if j-i+1 == len(s1) {
			if reflect.DeepEqual(s1Count, s2Count) {
				return true
			}
		}
		if j-i+1 < len(s1) {
			j++
		} else {
			s2Count[s2[i]-'a']--
			i++
			j++
		}
	}

	return false;
}


func main() {
	// INPUT
	s1 := "ab";
    s2 := "eidbaooo";

	// OUTPUT
	result := checkInclusion(s1, s2)
	fmt.Println(result)
}