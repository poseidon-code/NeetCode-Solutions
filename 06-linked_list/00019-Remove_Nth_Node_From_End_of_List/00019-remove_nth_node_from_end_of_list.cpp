// 19: Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/


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



class Solution {
public:
    // SOLUTION
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode *slow=head, *fast=head;
        for (int i=0; i<=n; i++) fast = fast->next;

        while (fast != NULL) {
            slow = slow->next;
            fast = fast->next;
        }

        slow->next = slow->next->next;
        return head;
    }
};



int main() {
    Solution o;
    ListNode* ll = NULL;

    // INPUT
    vector<int> li = {1,2,3,4,5};
    int n = 2;
    llInput(&ll, li);

    // OUTPUT
    auto result = o.removeNthFromEnd(ll, n);
    llOutput(result);

    return 0;
}
