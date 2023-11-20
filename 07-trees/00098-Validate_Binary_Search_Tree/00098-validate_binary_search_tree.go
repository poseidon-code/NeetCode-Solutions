// 98: Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/

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
func traverse(root, min, max *TreeNode) bool {
	if root==nil {return true}
	if min!=nil && root.val.(int)<=min.val.(int) {return false}
	if max!=nil && root.val.(int)>=max.val.(int) {return false}

	l := traverse(root.left, min, root)
	r := traverse(root.right, root, max)

	return l && r
}

func isValidBST(root *TreeNode) bool {
    return traverse(root, nil, nil)
}


func main() {
    tree := Tree{}

    // INPUT
    tn := []any{2,1,3}
    createTree(&tree.root, tn)

    // OUTPUT
    result := isValidBST(tree.root)
	fmt.Println(result)
}
