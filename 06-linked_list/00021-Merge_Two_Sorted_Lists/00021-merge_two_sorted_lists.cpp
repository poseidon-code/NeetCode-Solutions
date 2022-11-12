// 21: Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/


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
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        if (list1 == NULL) return list2;
        if (list2 == NULL) return list1;

        ListNode *p = list1;
        if (list1->val > list2->val) {
            p = list2;
            list2 = list2->next;
        } else {
            list1 = list1->next;
        }

        ListNode *c = p;
        while (list1 && list2) {
            if (list1->val < list2->val) {
                c->next = list1;
                list1 = list1->next;
            } else {
                c->next = list2;
                list2 = list2->next;
            }
            c = c->next;
        }

        if (!list1) c->next = list2;
        else c->next = list1;

        return p;
    }
};



int main() {
    Solution o;
    ListNode* head1 = NULL;
    ListNode* head2 = NULL;

    // INPUT
    vector<int> li1 = {1,2,4};
    vector<int> li2 = {1,3,4};
    llInput(&head1, li1);
    llInput(&head2, li2);

    // OUTPUT
    auto result = o.mergeTwoLists(head1, head2);
    llOutput(result);

    return 0;
}
