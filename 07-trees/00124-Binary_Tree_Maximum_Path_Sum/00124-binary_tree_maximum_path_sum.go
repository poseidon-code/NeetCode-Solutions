// 124: Binary Tree Maximum Path Sum
// https://leetcode.com/problems/binary-tree-maximum-path-sum

package main

import (
	"fmt"
	"math"
)

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
var max_sum = math.MinInt32

func traverse(root *TreeNode) int {
    if root==nil {return 0}
    c := func()int{if root.val==nil {return 0} else {return root.val.(int)}}()

    ls := traverse(root.left)
    rs := traverse(root.right)
    max_sum = int(math.Max(float64(max_sum), float64(ls+rs+c)))
    return int(math.Max(float64(ls), float64(rs))) + c
}

func maxPathSum(root *TreeNode) int {
    v := traverse(root)
    if max_sum==math.MinInt32 {return v}
    return max_sum
}

func main() {
    tree := Tree{}

    // INPUT
    tn := []any{-10,9,20,nil,nil,15,7}
    createTree(&tree.root, tn)

    // OUTPUT
    result := maxPathSum(tree.root)
    fmt.Println(result)
}