// 212: Word Search Ii
// https://leetcode.com/problems/word-search-ii/

package main

import (
	"fmt"
)


type TrieNode struct{
	val rune
    word string
    next map[rune]*TrieNode
}


func dfs(board [][]byte, res *[]string, node *TrieNode, i,j int, visited map[string]bool){
    k := fmt.Sprintf("%v_%v", i,j)

    if i<0 || j<0 || i>=len(board) || j>=len(board[0]) || visited[k] { return }
    nn, ok:= node.next[rune(board[i][j])]
    if !ok {return}
    
    if len(nn.word)>0 {
        *res = append(*res, nn.word)
        nn.word = ""
    }
    
    visited[k] = true
    dfs(board, res, nn, i+1, j, visited)
    dfs(board, res, nn, i-1, j, visited)
    dfs(board, res, nn, i, j+1, visited)
    dfs(board, res, nn, i, j-1, visited)
    visited[k] = false
}


// SOLUTION
func findWords(board [][]byte, words []string) []string {
    node := &TrieNode{next : map[rune]*TrieNode{}}
    for _, w := range words {
        p := node
        for i, b := range w {
            if _, ok := p.next[b]; !ok {
                p.next[b] = &TrieNode{val:b, next:map[rune]*TrieNode{}} 
            }
            if i==len(w)-1 {p.next[b].word = w}
            p = p.next[b]
        }
    }

    result := []string{}
    for i:=0; i<len(board); i++ {
        for j:=0; j<len(board[i]); j++ {
            dfs(board, &result, node, i,j, map[string]bool{})
        }
    }
    
    return result
}


func main() {
    // INPUT
    board := [][]byte{
        {'o','a','a','n'},
        {'e','t','a','e'},
        {'i','h','k','r'},
        {'i','f','l','v'},
    }
    words := []string{"oath","pea","eat","rain"}

    // OUTPUT
    fmt.Println(findWords(board, words))
}