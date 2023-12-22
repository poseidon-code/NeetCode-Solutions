// 211: Design Add And Search Words Data Structure
// https://leetcode.com/problems/design-add-and-search-words-data-structure/

package main

import "fmt"


type WordDictionary struct {
    children [26]*WordDictionary
	isEndOfWord bool
}


func Constructor() WordDictionary {
	return WordDictionary{}
}


func (this *WordDictionary) AddWord(word string)  {
    current := this

	for _, c := range word {
		if current.children[c - 'a'] == nil {
			current.children[c - 'a'] =  &WordDictionary{}
		}
		current = current.children[c - 'a'];
	}

	current.isEndOfWord = true;
}


func (this *WordDictionary) Search(word string) bool {
	current := this

	for i:=0; i<len(word); i++ {
		c := word[i]

		if c == '.' {
			for _, ch := range current.children {
				if ch!=nil && ch.Search(word[i+1:]) {return true}
			}
			return false
		}

		if (current.children[c - 'a'] == nil) {return false}
		current = current.children[c - 'a'];
	}

	return current!=nil && current.isEndOfWord
}


func main() {
	wd := Constructor()

	// OPERATIONS
	wd.AddWord("bad")
    wd.AddWord("dad")
    wd.AddWord("mad")
    fmt.Println(wd.Search("pad"))
    fmt.Println(wd.Search("bad"))
    fmt.Println(wd.Search(".ad"))
    fmt.Println(wd.Search("b.."))
}