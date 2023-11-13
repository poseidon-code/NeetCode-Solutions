// 1448: Count Good Nodes In Binary Tree
// https://leetcode.com/problems/count-good-nodes-in-binary-tree/


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
    int traverse(TreeNode *root, int max) {
        if(root!=NULL && root->val!=NULL && root->val>max)
            max = root->val;
        return root!=NULL && root->val!=NULL ? (root->val == max) + traverse(root->left, max) + traverse(root->right, max) : 0;
    }

public:
    // SOLUTION
    int goodNodes(TreeNode* root) {
        return traverse(root, -10000);
    }
};



int main() {
    Solution o;
    TreeNode* root = NULL;

    // INPUT
    vector<int> tn = {3,1,4,3,NULL,1,5};
    createTree(&root, tn);

    // OUTPUT
    auto result = o.goodNodes(root);
    cout << result << endl;

    return 0;
}
