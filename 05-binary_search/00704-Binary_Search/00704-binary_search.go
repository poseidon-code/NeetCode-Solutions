// 704: Binary Search
// https://leetcode.com/problems/binary-search/


package main

import "fmt"

// SOLUTION
func search(nums []int, target int) int {
	l, h := 0, len(nums) - 1

	for l <= h {
		m := l + (h-l)/2
		
		if nums[m] < target {
			l = m+1
		} else if nums[m] > target {
			h = m-1
		} else {
			return m
		}
	}

	return -1
}



func main() {
	// INPUT
	nums := []int{-1,0,3,5,9,12}
	target := 9

	// OUTPUT
	result := search(nums, target)
	fmt.Println(result)
}