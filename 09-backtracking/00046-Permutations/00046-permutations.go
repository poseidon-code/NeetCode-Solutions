// 46: Permutations
// https://leetcode.com/problems/permutations/

package main

import "fmt"


func backtrack(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		*result = append(*result, nums)
		return
	}

	for i:=start; i<len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]		
		backtrack(nums, (start + 1), result)
		nums[i], nums[start] = nums[start], nums[i]
	}
}


// SOLUTION
func permute(nums []int) [][]int {
	var result [][]int
	backtrack(nums, 0, &result)
	return result;
}

func main() {
    // INPUT
    nums := []int{0,1}

    // OUTPUT
    result := permute(nums)

	fmt.Println(result)
}
