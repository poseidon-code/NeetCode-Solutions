// 51: N Queens
// https://leetcode.com/problems/n-queens

package main

import (
	"fmt"
	"strings"
)

func backtrack(n int, row int, nQueens []string, result *[][]string) {
	if row == n {
		t := make([]string, len(nQueens))
		copy(t, nQueens)
		*result = append(*result, t)
		return
	}

	for col := 0; col < n; col++ {
		if isValid(nQueens, row, col, n) {
			r := []rune(nQueens[row])
			r[col] = 'Q'
			nQueens[row] = string(r)
			backtrack(n, row+1, nQueens, result)
			r = []rune(nQueens[row])
			r[col] = '.'
			nQueens[row] = string(r)
		}
	}
}

func isValid(nQueens []string, row, col, n int) bool {
	for i := 0; i < row; i++ {
		if nQueens[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if nQueens[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if nQueens[i][j] == 'Q' {
			return false
		}
	}

	return true
}

// SOLUTION
func solveNQueens(n int) [][]string {
	nQueens := make([]string, n)
	for i := 0; i < n; i++ {
		nQueens[i] = strings.Repeat(".", n)
	}
	var result [][]string
	backtrack(n, 0, nQueens, &result)
	return result
}

func main() {
	// INPUT
	n := 4

	// OUTPUT
	result := solveNQueens(n)

	fmt.Println(result)
}
