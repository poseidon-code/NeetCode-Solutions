// 739: Daily Temperstures
// https://leetcode.com/problems/daily-temperatures/

package main

import "fmt"

// SOLUTION
func dailyTemperatures(temperatures []int) []int {
	var s []int
	result := make([]int, len(temperatures))

	for i:=0; i<len(temperatures); i++ {
		for len(s)!=0 && temperatures[i]>temperatures[s[len(s)-1]] {
			result[s[len(s)-1]] = i - s[len(s)-1]
			s = s[:len(s)-1]
		}
		s = append(s, i);
	}

	return result;
}

func main() {
	// INPUT
	temperatures := []int{73,74,75,71,69,72,76,73}

	// OUTPUT
	result := dailyTemperatures(temperatures)
	fmt.Println(result)
}