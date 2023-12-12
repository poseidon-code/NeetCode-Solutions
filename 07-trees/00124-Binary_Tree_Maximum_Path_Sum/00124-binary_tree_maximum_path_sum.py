# 124: Binary Tree Maximum Path Sum
# https://leetcode.com/problems/binary-tree-maximum-path-sum



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


INT_MIN = -2**32

class Solution:
    # SOLUTION
    def __init__(self) -> None:
        self.max_sum = INT_MIN
        
    def maxPathSum (self, root: TreeNode) -> int:
        def traverse (root: TreeNode) -> int:
            if root==None : return 0
            c = 0 if root.val==None else root.val
            
            ls = max(traverse(root.left),0)
            rs = max(traverse(root.right),0)

            self.max_sum = max(self.max_sum, ls+rs+c)
            return max(ls, rs) + c
        
        v = traverse(root)
        if self.max_sum==INT_MIN : return v
        return self.max_sum




if __name__ == "__main__":
    o = Solution()
    tree = Tree()

    # INPUT
    tn = [-10,9,20,None,None,15,7]
    tree.createTree(tn)

    # OUTPUT
    result = o.maxPathSum(tree.root)
    print(result)
