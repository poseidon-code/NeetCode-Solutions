// 15: 3Sum
// https://leetcode.com/problems/3sum/

package main

import (
	"fmt"
	"sort"
)

// SOLUTION
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i:=0; i<len(nums); i++ {
		if i>0 && nums[i]==nums[i-1] {continue}

		l := i+1
		r := len(nums) - 1

		for l<r {
			s := nums[i] + nums[l] + nums[r]
			if s>0 {
				r--
			} else if (s<0) {
				l++
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for nums[l]==nums[l+1] {l++}
				for nums[r]==nums[r-1] {r--}
				l++
				r--
			}
		}
	}

	return result
}


func main() {
	// INPUT
	nums := []int{-1,0,1,2,-1,-4}

	// OUTPUT
	result := threeSum(nums)
	fmt.Println(result)
}