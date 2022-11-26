// 143: Reorder List
// https://leetcode.com/problems/reorder-list/


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
    ListNode* reorderList(ListNode* head) {
        if (head == NULL || head->next == NULL) return head;
        
        ListNode *prev=NULL, *slow=head, *fast=head, *l1=head, *l2=NULL;

        while (fast!=NULL && fast->next!=NULL) {
            prev = slow;
            slow = slow->next;
            fast = fast->next->next;
        }
        prev->next = NULL;

        ListNode *p=NULL, *c=slow, *n=NULL;
        while (c!=NULL) {
            n=c->next;
            c->next=p;
            p=c;
            c=n;
        }
        l2 = p;

        while (l1!=NULL) {
            ListNode *n1=l1->next, *n2=l2->next;
            l1->next = l2;
            if (n1==NULL) break;

            l2->next=n1;
            l1=n1;
            l2=n2;
        }

        return head;
    }
};



int main() {
    Solution o;
    ListNode* ll = NULL;

    // INPUT
    vector<int> li = {1,2,3,4};
    llInput(&ll, li);

    // OUTPUT
    auto result = o.reorderList(ll);
    llOutput(result);

    return 0;
}
