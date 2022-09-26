// 875: Koko Eating Bananas
// https://leetcode.com/problems/koko-eating-bananas/


package main

import "fmt"

// SOLUTION
func minEatingSpeed(piles []int, h int) int {
    l, r := 1, int(10e9)

	for l < r {
		total := 0
		m := (l + r) / 2

		for _, p := range piles {
			total += (p + m - 1) / m
		}
		if (total > h) {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func main() {
	// INPUT
	piles := []int{3,6,7,11}
	h := 8

	// OUTPUT
	result := minEatingSpeed(piles, h)
	fmt.Println(result)
}