// 79: Word Search
// https://leetcode.com/problems/word-search

package main

import (
	"fmt"
)

func backtrack(board [][]string, i, j int, word string) bool {
	if len(word)==0 {return true}
    if i<0 || i>=len(board) || j<0 || j>=len(board[0]) || board[i][j]!=string(word[0]) {return false}

    c := board[i][j]
    board[i][j] = "*"
    s := word[1:]
    
    result := backtrack(board, i-1, j, s) || backtrack(board, i+1, j, s) || backtrack(board, i, j-1, s) || backtrack(board, i, j+1, s)
    board[i][j] = c

    return result
}


// SOLUTION
func exist(board [][]string, word string) bool {
    for i:=0; i<len(board); i++ {
        for j:=0; j<len(board[0]); j++ {
            if backtrack(board, i, j, word) {
                return true
            }
        }
    }

    return false
}


func main() {
    // INPUT
    board := [][]string{
        {"A","B","C","E"},
        {"S","F","C","S"},
        {"A","D","E","E"},
    }
    word := "ABCCED"

    // OUTPUT
    fmt.Println(exist(board, word))
}
