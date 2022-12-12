// 141: Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/

package main

import "fmt"

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

func createLoop(head **ListNode, p int) {
    if (*head).next==nil {return}
	if p==-1 {return}

	tail, node := *head, *head
	for tail.next!=nil {tail = tail.next}
	for i:=0; i<p; i++ {node = node.next}

	tail.next = node
}

// SOLUTION
func hasCycle(head *ListNode) bool {
    if (head==nil) {return false}
    slow, fast := head, head
    
    for (fast.next!=nil && fast.next.next!=nil) {
        slow = slow.next
        fast = fast.next.next
        if (slow == fast) {return true}
    }

    return false
}

func main() {
    ll := LinkedList{}

    // INPUT
    li := []int{3,2,0,-4}
    llInput(&ll.head, li)
    createLoop(&ll.head, 1)

    // OUTPUT
    result := hasCycle(ll.head)
    fmt.Println(result)
}