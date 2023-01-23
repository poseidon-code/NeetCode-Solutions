// 110: Balanced Binary Tree
// https://leetcode.com/problems/balanced-binary-tree/

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
func depth(root *TreeNode) int {
    if root == nil {return 0}
    return int(math.Max(float64(depth(root.left)), float64(depth(root.right)))) + 1
}

func isBalanced(root *TreeNode) bool {
    if root == nil {return true}
    left := depth(root.left)
    right := depth(root.right)
    return math.Abs(float64(left - right)) <= 1 && isBalanced(root.left) && isBalanced(root.right)
    
}


func main() {
    tree := Tree{}

    // INPUT
    tn := []any{3,9,20,nil,nil,15,7}
    createTree(&tree.root, tn)

    // OUTPUT
    result := isBalanced(tree.root)
    fmt.Println(result)
}