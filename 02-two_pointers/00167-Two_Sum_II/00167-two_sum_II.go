// 167: Two Sum II
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

package main

import "fmt"

// SOLUTION
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers) - 1

	for l<r {
		s := numbers[l] + numbers[r]
		if s==target {
			return []int{l+1, r+1}
		} else if s<target {
			l++
		} else {
			r--
		}
	}
	
	return []int{}
}


func main() {
	// INPUT
	numbers := []int{2,7,11,15}
	target := 9

	// OUTPUT
	result := twoSum(numbers, target)
	fmt.Println(result)
}