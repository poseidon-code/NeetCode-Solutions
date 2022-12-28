// 23: Merge k Sorted Lists
// https://leetcode.com/problems/merge-k-sorted-lists

package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
    val int
    next *ListNode
}

type LinkedList struct {
    head *ListNode
}

func llInput(head **ListNode, inputs []int) {
    for i:=len(inputs)-1; i>=0; i-- {
        node := &ListNode{inputs[i], *head}
        *head = node
    }
}

func llOutput(head *ListNode) {
    for head != nil {
        fmt.Print(head.val, " -> ")
        head = head.next
    }
    fmt.Println("NULL")
}

type minHeap []*ListNode
func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *minHeap) Pop() interface{} {
    min := (*h)[len(*h) - 1]
    *h = (*h)[:len(*h) - 1]
    return min
}

// SOLUTION
func mergeKLists(lists []*ListNode) *ListNode {
    pq := &minHeap{}
	heap.Init(pq)
	var head, tail, next *ListNode

	for i:=0; i<len(lists); i++ {
		if lists[i]!=nil {
			heap.Push(pq, lists[i])
		}
	}

	for pq.Len() > 0 {
		if tail != nil {
			next = heap.Pop(pq).(*ListNode)
			if next.next != nil {
				heap.Push(pq, next.next)
			}
			tail.next = next
			tail = next
		} else {
			head = heap.Pop(pq).(*ListNode)
			tail = head
			if head.next != nil {
				heap.Push(pq, head.next)
			}
		}
	}

	return head
}


func main() {
    ll1, ll2, ll3 := LinkedList{}, LinkedList{}, LinkedList{}
    
    // INPUT
    li1, li2, li3 := []int{1,4,5}, []int{1,3,4}, []int{2,6}
    llInput(&ll1.head, li1)
    llInput(&ll2.head, li2)
    llInput(&ll3.head, li3)
    lists := []*ListNode{ll1.head, ll2.head, ll3.head}

    // OUTPUT
    result := mergeKLists(lists)
    llOutput(result)
}