// 239: Sliding Window Maximum
// https://leetcode.com/problems/sliding-window-maximum/

package main

import "fmt"

// Simple Deque Data Structure
type Deque struct {indices []int}
func (d *Deque) push(i int) {d.indices = append(d.indices, i)}
func (d *Deque) getFirst() int {return d.indices[0]}
func (d *Deque) popFirst() {d.indices = d.indices[1:]}
func (d *Deque) getLast() int {return d.indices[len(d.indices)-1]}
func (d *Deque) popLast() {d.indices = d.indices[:len(d.indices)-1]}
func (d *Deque) empty() bool {return 0 == len(d.indices)}


// SOLUTION
func maxSlidingWindow(nums []int, k int) []int {
	dq := &Deque{}
	result := make([]int, len(nums)-k+1)

	j := 0

	for i := range nums {
		if !dq.empty() && (dq.getFirst()<i-k+1) {
			dq.popFirst()
		}

		for !dq.empty() && nums[i] > nums[dq.getLast()] {
			dq.popLast()
		}

		dq.push(i)

		if i >= k-1 {
			result[j] = nums[dq.getFirst()]
			j++
		}
	}

	return result
}


func main() {
	// INPUT
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3

	// OUTPUT
	result := maxSlidingWindow(nums, k)
	fmt.Println(result)
}