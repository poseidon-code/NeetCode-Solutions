// 00543: Diameter of Binary Tree
// https://leetcode.com/problems/diameter-of-binary-tree/


import java.util.Queue;
import java.util.ArrayList;
import java.util.LinkedList;


class Solution {
    public static class TreeNode {
        Integer val;
        TreeNode left;
        TreeNode right;
        TreeNode() {}
        TreeNode(Integer val) { this.val = val; }
        TreeNode(Integer val, TreeNode left, TreeNode right) { this.val = val; this.left = left; this.right = right; }
    }

    public static class Tree {
        TreeNode root;
        Tree() { this.root = null; }
    }

    private static TreeNode creator(Integer[] values, TreeNode root, int i, int n) {
        if (n==0) return null;
        if (i<n) {
            TreeNode temp = new TreeNode(values[i]);
            root = temp;
            root.left = creator(values, root.left, 2*i+1, n);
            root.right = creator(values, root.right, 2*i+2, n);
        }
        return root;
    }

    public static void createTree(Tree tree, Integer[] inputs) {
        tree.root = creator(inputs, tree.root, 0, inputs.length);
    }

    public static void showTree(TreeNode root) {
        Queue<TreeNode> q = new LinkedList<>();
        ArrayList<ArrayList<Integer>> result = new ArrayList<>();
        ArrayList<Integer> c = new ArrayList<>();
        if (root == null) { System.out.println("Empty !"); return; }
        q.add(root);
        q.add(null);
        while (!q.isEmpty()) {
            TreeNode t = q.poll();
            if (t==null) {
                result.add(c);
                c = new ArrayList<>();
                if (q.size() > 0) q.add(null);
            } else {
                c.add(t.val);
                if (t.left!=null) q.add(t.left);
                if (t.right!=null) q.add(t.right);
            }
        }

        System.out.print("["); for (var x : result) {
            System.out.print("["); for (var y : x) {
                if (y==null) { System.out.print("NULL,"); continue; }
                System.out.print(y+",");
            } System.out.print("\b],");
        } System.out.println("\b]");
    }


    // SOLUTION
    int result = 0;

    private int dfs(TreeNode root) {
        if (root == null) return 0;

        int left = dfs(root.left);
        int right = dfs(root.right);

        result = Math.max(result, left + right);
        return 1 + Math.max(left, right);
    }

    public int diameterOfBinaryTree(TreeNode root) {
        if (root == null) return 0;
        dfs(root);
        return result;
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        Tree tree = new Tree();

        // INPUT
        Integer[] tn = {1,2,3,4,5};
        createTree(tree, tn);

        // OUTPUT
        var result = o.diameterOfBinaryTree(tree.root);
        System.out.println(result);
    }
}