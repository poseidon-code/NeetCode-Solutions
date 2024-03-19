// 1064: Last Stone Weight
// https://leetcode.com/problems/last-stone-weight

package main

import (
	"container/heap"
	"fmt"
)

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// SOLUTION
func lastStoneWeight(stones []int) int {
	pq := Heap(stones)
	heap.Init(&pq)

	for len(pq) > 1 {
		y := heap.Pop(&pq).(int)
		x := heap.Pop(&pq).(int)

		if x != y {
			heap.Push(&pq, y-x)
		}
	}

	if len(pq) == 0 {
		return 0
	}

	return pq[0]
}

func main() {
	// INPUT
	stones := []int{2, 7, 4, 1, 8, 1}

	// OUTPUT
	result := lastStoneWeight(stones)
	fmt.Println(result)
}
