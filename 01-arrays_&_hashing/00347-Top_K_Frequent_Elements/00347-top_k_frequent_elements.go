// 347: Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/

package main

import (
	"fmt"
	"sort"
)

// SOLUTION
func topKFrequent(nums []int, k int) []int {
	tmap := make(map[int]int)
	var buckets []int
	
    for _, v := range nums {
        tmap[v]++
    }

    for k := range tmap {
        buckets = append(buckets, k)
    }
    
    sort.Slice(buckets, func (i, j int) bool {
        return tmap[buckets[i]] > tmap[buckets[j]]
    })
    
    return buckets[:k]
}


func main() {
	// INPUT
	nums := []int{1,1,1,2,2,3}
	k := 2

	// OUTPUT
	result := topKFrequent(nums, k)
	fmt.Println(result)
}
