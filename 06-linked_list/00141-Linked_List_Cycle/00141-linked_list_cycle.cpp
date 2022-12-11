// 141: Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/


#include <iostream>
#include <vector>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

void llInput(ListNode** head, vector<int> inputs) {
    for (int i=inputs.size()-1; i>=0; i--) {
        ListNode* node = new ListNode(inputs[i], *head);
        *head = node;
    }
}

void llOutput(ListNode* head) {
    while (head != NULL) {
        cout << head->val << " -> ";
        head = head->next;
    }
    cout << "NULL" << endl;
}

void createLoop(ListNode** head, int p) {
    if (!(*head)->next) return;
    if (p==-1) return;

    ListNode *tail = *head, *node = *head;
    while (tail->next) tail = tail->next;
    for (int i=0; i<p; i++) node = node->next;
    
    tail->next = node;
}

class Solution {
public:
    // SOLUTION
    bool hasCycle(ListNode* head) {
        if (!head) return false;
        ListNode *slow=head, *fast=head;
        
        while (fast->next && fast->next->next) {
            slow = slow->next;
            fast = fast->next->next;
            if (slow == fast) return true;
        }
        
        return false;
    }
};



int main() {
    Solution o;
    ListNode* ll = NULL;

    // INPUT
    vector<int> li = {3,2,0,-4};
    llInput(&ll, li);
    createLoop(&ll, 1);

    // OUTPUT
    auto result = o.hasCycle(ll);
    cout << (result ? "true" : "false") << endl;

    return 0;
}
