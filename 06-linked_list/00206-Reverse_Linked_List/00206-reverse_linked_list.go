// 206: Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/

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
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.next == nil {return head}
    
    var previous *ListNode
    current := head
    next := head.next

    for current != nil {
        next = current.next;
        current.next = previous
        previous = current
        current = next
    }

    return previous
}

func main() {
    ll := LinkedList{}

    // INPUT
    li := []int{1,2,3,4,5}
    llInput(&ll.head, li)

    // OUTPUT
    result := reverseList(ll.head)
    llOutput(result)
}