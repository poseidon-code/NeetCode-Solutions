// 2: Add Two Numbers
// https://leetcode.com/problems/add-two-numbers/

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
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
    head := &ListNode{}
	t := head
	c := 0

	for c!=0 || l1!=nil || l2!=nil {
		if l1!=nil {
			c += l1.val
			l1 = l1.next
		}
		if l2!=nil {
			c += l2.val
			l2 = l2.next
		}

		t.next = &ListNode{val: c%10}
		t = t.next
		c /= 10
	}

	return head.next
}

func main() {
    ll1, ll2 := LinkedList{}, LinkedList{}
    
    // INPUT
    li1, li2 := []int{2,4,3}, []int{5,6,4}
    llInput(&ll1.head, li1)
    llInput(&ll2.head, li2)

    // OUTPUT
    result := addTwoNumbers(ll1.head, ll2.head)
    llOutput(result)
}