// 1: Two Sum
// https://leetcode.com/problems/two-sum/

package main

import "fmt"

// SOLUTION
func twoSum(nums []int, target int) []int {
	tmap := make(map[int]int)

	for i, n := range nums {
		_, m := tmap[n]

		if m {
			return []int{tmap[n], i}
		} else {
			tmap[target-n] = i
		}
	}

	return nil
}

func main() {
	// INPUT
	nums := []int{3,2,4}
	target := 6

	// OUTPUT
	result := twoSum(nums, target)
	fmt.Printf("[%d, %d]\n", result[0], result[1])
}
