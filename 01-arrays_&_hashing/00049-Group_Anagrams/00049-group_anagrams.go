// 49: Group Anagrams
// https://leetcode.com/problems/group-anagrams/

package main

import (
	"fmt"
	"sort"
	"strings"
)

// SOLUTION
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	tmap := make(map[string][]string, 0)
	
	for _, str := range strs {
		org := str
		ss := strings.Split(str, "")
		sort.Strings(ss)
		str = strings.Join(ss, "")
		tmap[str] = append(tmap[str], org)
	}

	for _, s := range tmap {
		result = append(result, s)
	}

	return result
}


func main() {
	// INPUT
	strs := []string{"eat","tea","tan","ate","nat","bat"}

	// OUTPUT
	result := groupAnagrams(strs)
	fmt.Println(result)
}
