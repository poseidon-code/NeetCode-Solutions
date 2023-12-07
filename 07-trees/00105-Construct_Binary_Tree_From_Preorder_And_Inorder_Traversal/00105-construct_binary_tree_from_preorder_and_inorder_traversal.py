# 105: Construct Binary Tree From Preorder And Inorder Traversal
# https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal

from typing import List

class TreeNode:
    def __init__(self, x: int):
        self.val = x
        self.left = None
        self.right = None

class Tree:
    def __init__(self) -> None:
        self.root = None

    def creator(self, values: list[int], root: TreeNode, i: int, n: int) -> TreeNode:
        if n==0 : return None
        if i<n:
            temp = TreeNode(values[i])
            root = temp
            root.left = self.creator(values, root.left, 2*i+1, n)
            root.right = self.creator(values, root.right, 2*i+2, n)
        return root

    def createTree(self, inputs: list[int]) -> None:
        self.root = self.creator(inputs, self.root, 0, len(inputs))
    
    def showTree(self, root: TreeNode) -> None:
        q = []
        result = [[]]
        c = []
        if root==None : return result
        q.append(root)
        q.append(None)
        while len(q)!=0:
            t = q.pop(0)
            if t==None:
                result.append(c)
                c = []
                if len(q) > 0 : q.append(None)
            else:
                c.append(t.val)
                if t.left!=None : q.append(t.left)
                if t.right!=None : q.append(t.right)

        print("[", end="")
        for x in result:
            if len(x)==0 : continue
            print("[", end="")
            for y in x:
                if y==None: 
                    print("NULL,", end="")
                    continue
                print(y, end=",")
            print("\b]", end=",")
        print("\b]")


class Solution:
    # SOLUTION
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        def build(preStart: int, inStart: int, inEnd: int, preorder: List[int], inorder: List[int]) -> TreeNode:
            if preStart > len(preorder) - 1 or inStart > inEnd: return None

            root = TreeNode(preorder[preStart])

            inIndex = 0
            for i in range(inStart, inEnd+1):
                if (inorder[i] == root.val):
                    inIndex = i

            root.left = build(preStart + 1, inStart, inIndex - 1, preorder, inorder)
            root.right = build(preStart + inIndex - inStart + 1, inIndex + 1, inEnd, preorder, inorder)

            return root

        return build(0, 0, len(inorder) - 1, preorder, inorder)


if __name__ == "__main__":
    o = Solution()
    tree = Tree()

    # INPUT
    preorder = [3,9,20,15,7]
    inorder = [9,3,15,20,7]

    # OUTPUT
    result = o.buildTree(preorder, inorder)
    tree.showTree(result)
