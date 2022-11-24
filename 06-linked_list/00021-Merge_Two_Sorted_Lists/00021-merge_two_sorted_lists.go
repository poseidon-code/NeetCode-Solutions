// 21: Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

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
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
    if list1 == nil {return list2}
    if list2 == nil {return list1}

    p := list1
    if list1.val > list2.val {
        p = list2
        list2 = list2.next
    } else {
        list1 = list1.next
    }

    c := p
    for list1!=nil && list2!=nil {
        if list1.val < list2.val {
            c.next = list1
            list1 = list1.next
        } else {
            c.next = list2
            list2 = list2.next
        }
        c = c.next
    }

    if list1 == nil {
        c.next = list2
    } else {
        c.next = list1
    }

    return p
}

func main() {
    ll1, ll2 := LinkedList{}, LinkedList{}
    
    // INPUT
    li1, li2 := []int{1,2,4}, []int{1,3,4}
    llInput(&ll1.head, li1)
    llInput(&ll2.head, li2)

    // OUTPUT
    result := mergeTwoLists(ll1.head, ll2.head)
    llOutput(result)
}