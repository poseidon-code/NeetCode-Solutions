// 206: Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/

package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
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
        fmt.Print(head.Val, " -> ")
        head = head.Next
    }
    fmt.Println("NULL")
}

// SOLUTION
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    
    var previous *ListNode
    current := head
    next := head.Next

    for current != nil {
        next = current.Next;
        current.Next = previous
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