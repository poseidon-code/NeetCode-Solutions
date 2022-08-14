// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func maxProfit(prices []int) int {
	minValue := prices[0]
	result := 0

	for i:=1; i<len(prices); i++ {
		minValue = int(math.Min(float64(minValue), float64(prices[i])))
        result = int(math.Max(float64(result), float64(prices[i] - minValue)))
	}

	return result
}


func main() {
	// INPUT
	prices := []int{7,1,5,3,6,4}

	// OUTPUT
	result := maxProfit(prices)
	fmt.Println(result)
}