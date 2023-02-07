// 572: Subtree Of Another Tree
// https://leetcode.com/problems/subtree-of-another-tree/

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

func isSame(root, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil {return true}
    if root == nil || subRoot == nil {return false}
    if root.val != subRoot.val {return false}
    return isSame(root.left, subRoot.left) && isSame(root.right, subRoot.right)
}

// SOLUTION
func isSubtree(root, subRoot *TreeNode) bool {
    if root == nil {return false}
    if isSame(root, subRoot) {return true}
    return isSubtree(root.left, subRoot) || isSubtree(root.right, subRoot)
}



func main() {
    tree, subTree := Tree{}, Tree{}

    // INPUT
    rn, srn := []any{3,4,5,1,2}, []any{4,1,2}
    createTree(&tree.root, rn)
    createTree(&subTree.root, srn)

    // OUTPUT
    result := isSubtree(tree.root, subTree.root)
    fmt.Println(result)
}