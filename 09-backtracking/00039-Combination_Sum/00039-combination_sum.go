// 39: Combination Sum
// https://leetcode.com/problems/combination-sum/

package main

import "fmt"


func backtrack(candidates []int, target int, start int, current []int, result *[][]int) {
	if start >= len(candidates) || target < 0 {return}

	if target == 0 {
		t := make([]int, len(current))
		copy(t, current)
		*result = append(*result, t)
		return
	}
	
	current = append(current, candidates[start])
	backtrack(candidates, target - candidates[start], start, current, result)
	current = current[:len(current)-1]
	backtrack(candidates, target, start + 1, current, result)
}


// SOLUTION
func combinationSum(candidates []int, target int) [][]int {
    var current []int
	var result [][]int
	backtrack(candidates, target, 0, current, &result)
	return result
}


func main() {
    // INPUT
    candidates := []int{2,3,6,7}
	target := 7

    // OUTPUT
    result := combinationSum(candidates, target)

	fmt.Println(result)
}
