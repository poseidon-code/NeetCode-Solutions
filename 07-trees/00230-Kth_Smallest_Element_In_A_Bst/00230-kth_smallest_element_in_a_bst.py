# 230: Kth Smallest Element In A Bst
# https://leetcode.com/problems/kth-smallest-element-in-a-bst/



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
    result, value = 0, 0
    
    def kthSmallest(self, root: TreeNode, k: int) -> int:
        def traverse(root: TreeNode) -> None:
            if not root or not root.val : return
            traverse(root.left)
            global value, result
            value-=1
            if value==0 : result = root.val
            traverse(root.right)
        
        global value
        value = k
        traverse(root)
        return result


if __name__ == "__main__":
    o = Solution()
    tree = Tree()

    # INPUT
    tn = [3,1,4,None,2]
    k = 1
    tree.createTree(tn)

    # OUTPUT
    result = o.kthSmallest(tree.root, k)
    print(result)
