// 199: Binary Tree Right Side View
// https://leetcode.com/problems/binary-tree-right-side-view/


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
    void traverse(TreeNode *root, int level, vector<int> &result) {
        if(root==NULL) return ;
        if(result.size()<level) result.push_back(root->val);
        traverse(root->right, level+1, result);
        traverse(root->left, level+1, result);
    }

public:
    // SOLUTION
    vector<int> rightSideView(TreeNode* root) {
        vector<int> result;
        traverse(root, 1, result);
        return result;
    }
};



int main() {
    Solution o;
    TreeNode* root = NULL;

    // INPUT
    vector<int> tn = {1,2,3,NULL,5,NULL,4};
    createTree(&root, tn);

    // OUTPUT
    auto result = o.rightSideView(root);
    cout<<"["; for (auto y : result) {
        if (!y) { cout << "NULL,"; continue; }
        cout<<y<<",";
    } cout<<"\b]";

    return 0;
}
