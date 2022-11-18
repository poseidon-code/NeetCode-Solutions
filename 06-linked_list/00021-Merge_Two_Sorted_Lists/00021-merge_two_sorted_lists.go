// 21: Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

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
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
    if list1 == nil {return list2}
    if list2 == nil {return list1}

    p := list1
    if list1.Val > list2.Val {
        p = list2
        list2 = list2.Next
    } else {
        list1 = list1.Next
    }

    c := p
    for list1!=nil && list2!=nil {
        if list1.Val < list2.Val {
            c.Next = list1
            list1 = list1.Next
        } else {
            c.Next = list2
            list2 = list2.Next
        }
        c = c.Next
    }

    if list1 == nil {
        c.Next = list2
    } else {
        c.Next = list1
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