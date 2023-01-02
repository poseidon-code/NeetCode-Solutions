// 25: Reverse Nodes in K-Group
// https://leetcode.com/problems/reverse-nodes-in-k-group/

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
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{}
    dummy.next = head

    previous := dummy
    current := dummy.next
    t := &ListNode{}

    count := k

    for current != nil {
        if count > 1 {
            t = previous.next
            previous.next = current.next
            current.next = current.next.next
            previous.next.next = t
            count--
        } else {
            previous = current
            current = current.next
            count = k

            end := current
            for i := 0; i < k; i++ {
                if end == nil {return dummy.next}
                end = end.next
            }
        }
    }

    return dummy.next
}

func main() {
    ll := LinkedList{}

    // INPUT
    li := []int{1,2,3,4,5}
    k := 3
    llInput(&ll.head, li)

    // OUTPUT
    result := reverseKGroup(ll.head, k)
    llOutput(result)
}