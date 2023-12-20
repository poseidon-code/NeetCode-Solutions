// 208: Implement Trie Prefix Tree
// https://leetcode.com/problems/implement-trie-prefix-tree

package main

import (
	"fmt"
)


type Trie struct {
    children [26]*Trie
	isWord bool
}


func Constructor() Trie {
	return Trie{}
}


func (this *Trie) Insert(word string)  {
    node := this
	current := 0
	
	for i:=0; i<len(word); i++ {
		current = int(word[i] - 'a')
		if node.children[current] == nil {
			node.children[current] = &Trie{}
		}
		node = node.children[current]
	}

	node.isWord = true
}


func (this *Trie) Search(word string) bool {
    node := this
	current := 0

	for i:=0; i<len(word); i++ {
		current = int(word[i] - 'a')
		if node.children[current] == nil {
			return false
		}
		node = node.children[current]
	}

	return node.isWord
}


func (this *Trie) StartsWith(prefix string) bool {
	node := this
	current := 0

	for i:=0; i<len(prefix); i++ {
		current = int(prefix[i] - 'a')
		if node.children[current] == nil {
			return false
		}
		node = node.children[current]
	}

	return true
}


func main() {
	trie := Constructor()

	// OPERATIONS
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}