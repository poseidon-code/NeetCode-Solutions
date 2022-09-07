// 84: Largest Rectangle in Histogram
// https://leetcode.com/problems/largest-rectangle-in-histogram/

package main

import (
	"fmt"
	"math"
)


// SOLUTION
func largestRectangleArea(heights []int) int {
	var stack []int
	l := len(heights)
	result := 0

	for i:=0; i<=l; i++ {
		var h int
		if i==l {h=0} else {h=heights[i]}

		if len(stack)==0 || h>=heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			tp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var t int
			if len(stack)==0 {t=i} else {t=i-1-stack[len(stack)-1]}
			result = int(math.Max(float64(result), float64(heights[tp] * t)))
			i--
		}
	}

	return result
}


func main() {
	// INPUT
	heights := []int{2,1,5,6,2,3}

	// OUTPUT
	result := largestRectangleArea(heights)
	fmt.Println(result)
}