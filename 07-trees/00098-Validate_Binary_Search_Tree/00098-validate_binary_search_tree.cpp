// 98: Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/



#include <iostream>
#include <queue>
#include <vector>
#include <stack>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

TreeNode* creator(vector<int> values, TreeNode** root, int i, int n) {
    if (n==0) return NULL;
    if (i<n) {
        TreeNode *temp = new TreeNode(values[i]);
        *root = temp;
        (*root)->left = creator(values, &((*root)->left), 2*i+1, n);
        (*root)->right = creator(values, &((*root)->right), 2*i+2, n);
    }
    return *root;
};

void createTree(TreeNode** root, vector<int> inputs) {
    creator(inputs, root, 0, inputs.size());
}

void showTree(TreeNode* root) {
    queue<TreeNode*> q;
    vector<vector<int>> result = {};
    vector<int> c;
    if (root==NULL) { cout << "Empty !" << endl; return; }
    q.push(root);
    q.push(NULL);
    while (!q.empty()) {
        TreeNode *t = q.front();
        q.pop();
        if (t==NULL) {
            result.push_back(c);
            c.resize(0);
            if (q.size() > 0) q.push(NULL);
        } else {
            c.push_back(t->val);
            if (t->left) q.push(t->left);
            if (t->right) q.push(t->right);
        }
    }

    cout<<"["; for (auto x : result) {
        cout<<"["; for (auto y : x) {
            if (!y) { cout << "NULL,"; continue; }
            cout<<y<<",";
        } cout<<"\b],";
    } cout<<"\b]"<<endl;
}



class Solution {
private:
    bool traverse(TreeNode *root, TreeNode* min, TreeNode* max) {
        if (root==NULL) return true;
        if (min && root->val<=min->val) return false;
        if (max && root->val>=max->val) return false;

        bool l = traverse(root->left, min, root);
        bool r = traverse(root->right, root, max);

        return l && r;
    }

public:
    // SOLUTION
    bool isValidBST(TreeNode* root) {
        return traverse(root, NULL, NULL);
    }
};



int main() {
    Solution o;
    TreeNode* root = NULL;

    // INPUT
    vector<int> tn = {2,1,3};
    createTree(&root, tn);

    // OUTPUT
    auto result = o.isValidBST(root);
    cout << (result ? "true" : "false") << endl;

    return 0;
}
