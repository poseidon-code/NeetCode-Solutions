// 36: Valid Sudoku
// https://leetcode.com/problems/valid-sudoku/

package main

import (
	"fmt"
	"strconv"
)

// SOLUTION
func isValidSudoku(board [][]string) bool {
	used1, used2, used3 := make([][]int, 9), make([][]int, 9), make([][]int, 9)
	for i := range used1 {
		used1[i], used2[i], used3[i] = make([]int, 9), make([]int, 9), make([]int, 9)
	}

	for i:=0; i<len(board); i++ {
		for j:=0; j<len(board[i]); j++ {
			if board[i][j]!="." {
				s, _ := strconv.Atoi(board[i][j])
				num := s - 1
				k := i/3 * 3 + j/3
				
				if used1[i][num]!=0 || used2[j][num]!=0 || used3[k][num]!=0 {
					return false
				}
				
				used1[i][num], used2[j][num], used3[k][num] = 1, 1, 1
			}
		}
	}

	return true
}

func main() {
	// INPUT
	board := [][]string{
		{"5","3",".",".","7",".",".",".","."},
        {"6",".",".","1","9","5",".",".","."},
        {".","9","8",".",".",".",".","6","."},
        {"8",".",".",".","6",".",".",".","3"},
        {"4",".",".","8",".","3",".",".","1"},
        {"7",".",".",".","2",".",".",".","6"},
        {".","6",".",".",".",".","2","8","."},
        {".",".",".","4","1","9",".",".","5"},
        {".",".",".",".","8",".",".","7","9"},
	}

	// OUTPUT
	result := isValidSudoku(board)
	fmt.Println(result)
}