// 00128: Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func longestConsecutive(nums []int) int {
	result := 0
	tmap := make(map[int]int)

	for _, i := range nums {
		if tmap[i]!=0 {continue}
		
		left := tmap[i-1]
		right := tmap[i+1]
		sum := left + right + 1
		result = int(math.Max(float64(result), float64(sum)))

		if left>0 {tmap[i-left] = sum}
		if right>0 {tmap[i+right] = sum}
		tmap[i] = sum		
	}

	return result
}


func main() {
	// INPUT
	nums := []int{100,4,200,1,3,2}

	// OUTPUT
	result := longestConsecutive(nums)
	fmt.Println(result)
}