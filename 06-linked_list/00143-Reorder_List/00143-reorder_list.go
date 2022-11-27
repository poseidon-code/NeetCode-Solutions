// 143: Reorder List
// https://leetcode.com/problems/reorder-list/

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
func reorderList(head *ListNode) *ListNode {
    if head == nil || head.next == nil {return head}

    var prev, slow, fast, l1, l2 *ListNode
    prev, slow, fast, l1, l2 = nil, head, head, head, nil

    for fast!=nil && fast.next!=nil {
        prev = slow
        slow = slow.next
        fast = fast.next.next
    }
    prev.next = nil

    var p, c, n *ListNode
    p, c, n = nil, slow, nil
    for c!=nil {
        n=c.next
        c.next=p
        p=c
        c=n
    }
    l2 = p

    for l1!=nil {
        n1, n2 := l1.next, l2.next
        l1.next = l2
        if n1==nil {break}

        l2.next=n1
        l1=n1
        l2=n2
    }

    return head
}

func main() {
    ll := LinkedList{}

    // INPUT
    li := []int{1,2,3,4}
    llInput(&ll.head, li)

    // OUTPUT
    result := reorderList(ll.head)
    llOutput(result)
}