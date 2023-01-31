// 100: Same Tree
// https://leetcode.com/problems/same-tree/

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
public:
    // SOLUTION
    bool isSameTree(TreeNode* p, TreeNode* q) {
        if (p == NULL && q == NULL) return true;
        if (p == NULL || q == NULL) return false;
        if (p->val != q->val) return false;
        return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
    }
};



int main() {
    Solution o;
    TreeNode *p = NULL, *q = NULL;

    // INPUT
    vector<int> pn = {1,2,3};
    createTree(&p, pn);
    vector<int> qn = {1,2,3};
    createTree(&q, qn);

    // OUTPUT
    auto result = o.isSameTree(p, q);
    cout << (result ? "true" : "false") << endl;

    return 0;
}
