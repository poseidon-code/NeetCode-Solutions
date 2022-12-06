// 2: Add Two Numbers
// https://leetcode.com/problems/add-two-numbers/

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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *head = new ListNode;
        ListNode *t = head;
        int c = 0;

        while (c || l1 || l2) {
            c += (l1 ? l1->val : 0) + (l2 ? l2->val : 0);
            
            ListNode *temp = new ListNode;
            temp->val = c%10;
            temp->next = NULL;
            
            t->next = temp;
            t = t->next;
            c /= 10;

            if (l1) l1 = l1->next;
            if (l2) l2 = l2->next;
        }
        
        return head->next;
    }
};



int main() {
    Solution o;
    ListNode* ll1 = NULL;
    ListNode* ll2 = NULL;

    // INPUT
    vector<int> li1 = {2,4,3};
    vector<int> li2 = {5,6,4};
    llInput(&ll1, li1);
    llInput(&ll2, li2);

    // OUTPUT
    auto result = o.addTwoNumbers(ll1, ll2);
    llOutput(result);

    return 0;
}
