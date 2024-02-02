// 90: Subsets Ii
// https://leetcode.com/problems/subsets-ii

package main

import "fmt"


func backtrack(nums []int, start int, current []int, result *[][]int) {
	*result = append(*result, current)

	for i:=start; i<len(nums); i++ {
		if (i>start && nums[i] == nums[i-1]) {continue}
		current = append(current, nums[i])
		backtrack(nums, (i + 1), current, result)
		current = current[:len(current)-1]
	}
}


// SOLUTION
func subsetsWithDup(nums []int) [][]int {
    var current []int
	var result [][]int
	backtrack(nums, 0, current, &result)
	return result;
}

func main() {
    // INPUT
    nums := []int{1,2,2}

    // OUTPUT
    result := subsetsWithDup(nums)

	fmt.Println(result)
}
