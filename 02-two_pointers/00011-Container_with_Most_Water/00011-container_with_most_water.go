// 11: Container with Most Water
// https://leetcode.com/problems/container-with-most-water/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	result := 0

	for l<r {
		result = int(math.Max(float64(result), float64((r-l))*math.Min(float64(height[l]), float64(height[r]))))
		if height[l]<height[r] {
			l++
		} else {
			r--
		}
	}

	return result
}


func main() {
	// INPUT
	height := []int{1,8,6,2,5,4,8,3,7}

	// OUTPUT
	result := maxArea(height)
	fmt.Println(result)
}