// 155: Min stack
// https://leetcode.com/problems/min-stack/

package main

import "fmt"

// SOLUTION
type MinStack struct {
    s1, s2 []int
}

func (this *MinStack) push(x int) {
    this.s1 = append(this.s1, x)
    if len(this.s2)==0 || x <= this.getMin() {
        this.s2 = append(this.s2, x)
    }
}

func (this *MinStack) pop() {
    if this.s1[len(this.s1)-1] == this.getMin() {
        this.s2 = this.s2[:len(this.s2)-1]
    }
    this.s1 = this.s1[:len(this.s1)-1]
}

func (this *MinStack) top() int {
    return this.s1[len(this.s1)-1]
}

func (this *MinStack) getMin() int {
    return this.s2[len(this.s2)-1]
}


func main() {
    var o MinStack

    // OPERATIONS
    o.push(-2)
    o.push(0)
    o.push(-3)
    fmt.Print(o.getMin(), " ")
    o.pop()
    o.top()
    fmt.Print(o.getMin(), " ")

    fmt.Println()
}