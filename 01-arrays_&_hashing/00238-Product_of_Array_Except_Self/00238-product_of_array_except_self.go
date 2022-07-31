// 238 : Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/

package main

import "fmt"

// SOLUTION
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for i := range result {result[i] = 1}

	for i, suf, pre, n:=0, 1, 1, len(nums); i<n; i++ {
		result[i] *= pre;
		pre *= nums[i];
		result[n-1-i] *= suf;
		suf *= nums[n-1-i];
	}

	return result
}


func main() {
	// INPUT
	nums := []int{1,2,3,4}

	// OUTPUT
	result := productExceptSelf(nums)
	fmt.Println(result)
}