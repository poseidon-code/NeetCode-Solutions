// 74: Search a 2D Matrix
// https://leetcode.com/problems/search-a-2d-matrix/

package main

import "fmt"

// SOLUTION
func searchMatrix(matrix [][]int, target int) bool {
    rows := len(matrix)
	cols := len(matrix[0])
	row := 0
	col := cols-1
		
	for row<rows && col>-1 {
		cur := matrix[row][col]
		if cur == target {
			return true
		}
		if target > cur {
			row++
		} else {
			col--
		}
	}
	
	return false
}

func main() {
	// INPUT
	matrix := [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}
	target := 3

	// OUTPUT
	result := searchMatrix(matrix, target)
	fmt.Println(result)
}