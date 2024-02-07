// 40: Combination Sum Ii
// https://leetcode.com/problems/combination-sum-ii

package main

import (
	"fmt"
	"sort"
)


func backtrack(candidates []int, target int, sum int, start int, current []int, result *[][]int) {
	if sum > target {return}

	if sum == target {
		t := make([]int, len(current))
		copy(t, current)
		*result = append(*result, t)
		return
	}
	
	for i:=start; i<len(candidates); i++ {
		if i>start && candidates[i] == candidates[i-1] {continue}
		current = append(current, candidates[i])
		backtrack(candidates, target, sum + candidates[i], i+1, current, result)
		current = current[:len(current)-1]
	}
}


// SOLUTION
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
    var current []int
	var result [][]int
	backtrack(candidates, target, 0, 0, current, &result)
	return result
}


func main() {
    // INPUT
    candidates := []int{2,5,2,1,2}
	target := 5

    // OUTPUT
    result := combinationSum2(candidates, target)

	fmt.Println(result)
}
