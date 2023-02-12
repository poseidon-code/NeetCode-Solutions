// 235: Lowest Common Ancestor Of A Binary Search Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if (p.val.(int) < root.val.(int) && q.val.(int) < root.val.(int)) {
        return lowestCommonAncestor(root.left, p, q);
    } else if (p.val.(int) > root.val.(int) && q.val.(int) > root.val.(int)) {
        return lowestCommonAncestor(root.right, p, q);
    } else {
        return root;
    }
}



func main() {
    tree := Tree{}

    // INPUT
    tn := []any{6,2,8,0,4,7,9,nil,nil,3,5}
    createTree(&tree.root, tn)
    p := TreeNode{2, nil, nil}
    q := TreeNode{8, nil, nil}

    // OUTPUT
    result := lowestCommonAncestor(tree.root, &p, &q)
    fmt.Println(result.val)
}