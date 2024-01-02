// 78: Subsets
// https://leetcode.com/problems/subsets/

package main

import "fmt"


func backtrack(nums []int, start int, current []int, result *[][]int) {
	*result = append(*result, current)

	for i:=start; i<len(nums); i++ {
		current = append(current, nums[i])
		backtrack(nums, (i + 1), current, result);
		current = current[:len(current)-1]
	}
}


// SOLUTION
func subsets(nums []int) [][]int {
    var current []int
	var result [][]int
	backtrack(nums, 0, current, &result);
	return result;
}

func main() {
    // INPUT
    nums := []int{1,2,3}

    // OUTPUT
    result := subsets(nums);

	fmt.Println(result)
}
