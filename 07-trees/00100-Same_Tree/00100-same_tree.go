// 100: Same Tree
// https://leetcode.com/problems/same-tree/

package main

import "fmt"

type TreeNode struct {
    val any
    left *TreeNode
    right *TreeNode
}

type Tree struct {
    root *TreeNode
}

func creator(values []any, root **TreeNode, i, n int) *TreeNode {
    if n==0 {return nil}
    if i<n {
        temp := &TreeNode{values[i], nil, nil}
        *root = temp
        (*root).left = creator(values, &((*root).left), 2*i+1, n);
        (*root).right = creator(values, &((*root).right), 2*i+2, n);
    }
    return *root;
}

func createTree(root **TreeNode, inputs []any) {
    creator(inputs, root, 0, len(inputs))
}

func showTree(root *TreeNode) {
    var q []*TreeNode;
    var result [][]interface{}
    var c []interface{}
    if root==nil { fmt.Println("Empty !"); return; }
    q = append(q, root)
    q = append(q, nil)
    for len(q)!=0 {
        t := q[0]
        q = q[1:]
        if t==nil {
            result = append(result, c)
            c = make([]interface{}, 0)
            if len(q) > 0 {q = append(q, nil)}
        } else {
            c = append(c, t.val)
            if t.left!=nil {q = append(q, t.left)}
            if t.right!=nil {q = append(q, t.right)}
        }
    }

    fmt.Print("["); for _, x := range result {
        fmt.Print("["); for _, y := range x {
            if y==nil { fmt.Print("NULL,"); continue; }
            fmt.Print(y,",")
        }; fmt.Print("\b],")
    }; fmt.Println("\b]")
}

// SOLUTION
func isSameTree(p, q *TreeNode) bool {
    if p == nil && q == nil {return true}
    if p == nil || q == nil {return false}
    if p.val != q.val {return false}
    return isSameTree(p.left, q.left) && isSameTree(p.right, q.right)
}



func main() {
    p, q := Tree{}, Tree{}

    // INPUT
    pn, qn := []any{1,2,3}, []any{1,2,3}
    createTree(&p.root, pn)
    createTree(&q.root, qn)

    // OUTPUT
    result := isSameTree(p.root, q.root)
    fmt.Println(result)
}