// 25: Reverse Nodes in K-Group
// https://leetcode.com/problems/reverse-nodes-in-k-group/


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
    ListNode* reverseKGroup(ListNode* head, int k) {
        ListNode* dummy = new ListNode();
        dummy->next = head;

        ListNode* previous = dummy;
        ListNode* current = dummy->next;
        ListNode* t = NULL;

        int count = k;

        while (current != NULL) {
            if (count > 1) {
                t = previous->next;
                previous->next = current->next;
                current->next = current->next->next;
                previous->next->next = t;
                count--;
            } else {
                previous = current;
                current = current->next;
                count = k;

                ListNode* end = current;
                for (int i = 0; i < k; i++) {
                    if (end == NULL) return dummy->next;
                    end = end->next;
                }
            }
        }

        return dummy->next;
    }
};



int main() {
    Solution o;
    ListNode* ll = NULL;

    // INPUT
    vector<int> li = {1,2,3,4,5};
    int k = 3;
    llInput(&ll, li);

    // OUTPUT
    auto result = o.reverseKGroup(ll, k);
    llOutput(result);

    return 0;
}
