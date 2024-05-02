// 295: Find Median From Data Stream
// https://leetcode.com/problems/find-median-from-data-stream

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
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// SOLUTION
type MedianFinder struct {
	lower  Heap
	higher Heap
}

func Constructor() MedianFinder {
	return MedianFinder{lower: Heap{}, higher: Heap{}}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.higher) > 0 && num > this.higher[0] {
		heap.Push(&this.higher, num)
	} else {
		heap.Push(&this.lower, -1*num)
	}

	if len(this.lower) > len(this.higher)+1 {
		val := -1 * heap.Pop(&this.lower).(int)
		heap.Push(&this.higher, val)
	}

	if len(this.higher) > len(this.lower)+1 {
		val := heap.Pop(&this.higher).(int)
		heap.Push(&this.lower, -1*val)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.lower) > len(this.higher) {
		return float64(-1 * this.lower[0])
	} else if len(this.higher) > len(this.lower) {
		return float64(this.higher[0])
	}

	return float64(-1*this.lower[0]+this.higher[0]) / 2
}

func main() {
	o := Constructor()

	// OPERATIONS
	o.AddNum(1)
	o.AddNum(2)
	fmt.Println(o.FindMedian())
	o.AddNum(3)
	fmt.Println(o.FindMedian())
}
