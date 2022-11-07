// 206: Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/


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
    ListNode* reverseList(ListNode* head) {
        if (head == NULL || head->next == NULL) return head;
        
        ListNode* previous = NULL;
        ListNode* current = head;
        ListNode* next = current->next;

        while (current != NULL) {
            next = current->next;
            current->next = previous;
            previous = current;
            current = next;
        }

        return previous;
    }
};



int main() {
    Solution o;
    ListNode* head = NULL;

    // INPUT
    vector<int> li = {1,2,3,4,5};
    llInput(&head, li);

    // OUTPUT
    auto result = o.reverseList(head);
    llOutput(result);

    return 0;
}
