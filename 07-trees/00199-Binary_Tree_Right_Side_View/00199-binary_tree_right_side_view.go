// 199: Binary Tree Right Side View
// https://leetcode.com/problems/binary-tree-right-side-view/

package main

import (
	"fmt"
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
func traverse(root *TreeNode, level int, result *[]any) {
    if root==nil {return}
	if len(*result)<level {*result = append(*result, root.val)};
	traverse(root.right, level+1, result);
	traverse(root.left, level+1, result);
}

func rightSideView(root *TreeNode) []any {
	var result []any
    traverse(root, 1, &result);
    return result;
}


func main() {
    tree := Tree{}

    // INPUT
    tn := []any{1,2,3,nil,5,nil,4}
    createTree(&tree.root, tn)

    // OUTPUT
    result := rightSideView(tree.root)
    fmt.Print("["); for _, x := range result {
        if x==nil { fmt.Print("NULL,"); continue; }
        fmt.Print(x,",")
    }; fmt.Print("\b]")
}