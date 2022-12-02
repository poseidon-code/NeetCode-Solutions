// 19: Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

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

// SOLUTION
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    slow, fast := head, head
    for i:=0; i<=n; i++ {fast = fast.next}

    for fast != nil {
        slow = slow.next
        fast = fast.next
    }

    slow.next = slow.next.next
    return head
}

func main() {
    ll := LinkedList{}

    // INPUT
    li := []int{1,2,3,4,5}
    n := 2
    llInput(&ll.head, li)

    // OUTPUT
    result := removeNthFromEnd(ll.head, n)
    llOutput(result)
}