// 973: K Closest Points To Origin
// https://leetcode.com/problems/k-closest-points-to-origin

package main

import (
	"container/heap"
	"fmt"
	"math"
)

type PointEntry struct {
	dist int
	x    int
	y    int
}

type Heap []*PointEntry

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(*PointEntry)) }
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

const X = 0
const Y = 1

// SOLUTION
func kClosest(points [][]int, k int) [][]int {
	triples := make([]*PointEntry, 0, len(points))

	for _, point := range points {
		dist := math.Pow(math.Abs(float64(point[X])), 2) + math.Pow(math.Abs(float64(point[Y])), 2)
		triples = append(triples, &PointEntry{dist: int(dist), x: point[X], y: point[Y]})
	}

	result := make([][]int, 0)
	pq := Heap(triples)
	heap.Init(&pq)

	for i := 0; i < k; i++ {
		point := heap.Pop(&pq).(*PointEntry)
		result = append(result, []int{point.x, point.y})
	}

	return result
}

func main() {
	// INPUT
	points := [][]int{{1, 3}, {-2, 2}}
	k := 1

	// OUTPUT
	result := kClosest(points, k)
	fmt.Println(result)
}
