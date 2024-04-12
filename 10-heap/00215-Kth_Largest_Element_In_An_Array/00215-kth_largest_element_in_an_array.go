// 215: Kth Largest Element In An Array
// https://leetcode.com/problems/kth-largest-element-in-an-array

package main

import (
	"container/heap"
	"fmt"
)

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	n := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return n
}

// SOLUTION
func findKthLargest(nums []int, k int) int {
	pq := Heap(nums)
	heap.Init(&pq)

	for pq.Len() > k {
		heap.Pop(&pq)
	}

	return heap.Pop(&pq).(int)
}

func main() {
	// INPUT
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2

	// OUTPUT
	result := findKthLargest(nums, k)
	fmt.Println(result)
}
