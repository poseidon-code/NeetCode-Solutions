// 217: Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/

package main

import "fmt"

// SOLUTION
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		} else {
			m[n] = true
		}
	}

	return false
}

func main() {
	// INPUT
	nums := []int{1,2,3,1}

	// OUTPUT
	result := containsDuplicate(nums)
	fmt.Println(result)
}
