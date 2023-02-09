# 572: Subtree Of Another Tree
# https://leetcode.com/problems/subtree-of-another-tree/



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
    def isSame(self, p: TreeNode, q: TreeNode) -> bool:
        if p == None and q == None : return True
        if p == None or q == None : return False
        if p.val != q.val : return False
        return self.isSame(p.left, q.left) and self.isSame(p.right, q.right)

    # SOLUTION
    def isSubtree(self, root: TreeNode, subRoot: TreeNode) -> bool:
        if root == None : return False
        if self.isSame(root, subRoot) : return True
        return self.isSubtree(root.left, subRoot) or self.isSubtree(root.right, subRoot)



if __name__ == "__main__":
    o = Solution()
    tree, subTree = Tree(), Tree()

    # INPUT
    rn, srn = [3,4,5,1,2], [4,1,2]
    tree.createTree(rn)
    subTree.createTree(srn)

    # OUTPUT
    result = o.isSubtree(tree.root, subTree.root)
    print(result)
