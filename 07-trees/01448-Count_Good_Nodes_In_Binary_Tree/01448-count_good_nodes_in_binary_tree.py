# 1448: Count Good Nodes In Binary Tree
# https://leetcode.com/problems/count-good-nodes-in-binary-tree/



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
        result = []
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
    def goodNodes(self, root: TreeNode) -> int:        
        def traverse(root: TreeNode, max: int) -> None:
            if root!=None and root.val!=None and root.val>max:
                max = root.val
            return (root.val == max) + traverse(root.left, max) + traverse(root.right, max) if root!=None and root.val!=None else 0

        return traverse(root, -10000)


if __name__ == "__main__":
    o = Solution()
    tree = Tree()

    # INPUT
    tn = [3,1,4,3,None,1,5]
    tree.createTree(tn)

    # OUTPUT
    result = o.goodNodes(tree.root)
    print(result)
