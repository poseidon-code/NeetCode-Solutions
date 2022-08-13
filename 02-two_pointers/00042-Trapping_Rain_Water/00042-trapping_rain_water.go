// 42: Trapping Rain Water
// https://leetcode.com/problems/trapping-rain-water/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func trap(height []int) int {
	l, r := 0, len(height) - 1
	maxLeft, maxRight := height[l], height[r]
	result := 0

	for l<r {
		if maxLeft<=maxRight {
			l++
			maxLeft = int(math.Max(float64(maxLeft), float64(height[l])))
			result += maxLeft - height[l]
		} else {
			r--
			maxRight = int(math.Max(float64(maxRight), float64(height[r])))
			result += maxRight - height[r]
		}
	}

	return result
}


func main() {
	// INPUT
	height := []int{4,2,0,3,2,5}

	// OUTPUT
	result := trap(height)
	fmt.Println(result)
}