// 105: Construct Binary Tree From Preorder And Inorder Traversal
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal

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
func build(preStart int, inStart int, inEnd int, preorder []any, inorder []any) *TreeNode {
	if (preStart > len(preorder) - 1) || (inStart > inEnd) { return nil }
        
	root := &TreeNode{preorder[preStart], nil, nil}

	inIndex := 0
	for i := inStart; i <= inEnd; i++ {
		if (inorder[i] == root.val) {
			inIndex = i
		}
	}

	root.left = build(preStart + 1, inStart, inIndex - 1, preorder, inorder);
	root.right = build(preStart + inIndex - inStart + 1, inIndex + 1, inEnd, preorder, inorder);
	
	return root;
}

func buildTree(preorder, inorder []any) *TreeNode {
	return build(0, 0, len(inorder) - 1, preorder, inorder)
}


func main() {
    // INPUT
    preorder := []any{3,9,20,15,7}
    inorder := []any{9,3,15,20,7}

    // OUTPUT
    result := buildTree(preorder, inorder)
	showTree(result)
}
