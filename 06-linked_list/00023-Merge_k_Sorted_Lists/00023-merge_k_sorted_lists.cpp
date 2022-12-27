// 23: Merge k Sorted Lists
// https://leetcode.com/problems/merge-k-sorted-lists


#include <iostream>
#include <vector>
#include <queue>

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
private:
    struct compare {
        bool operator()(const ListNode* l, const ListNode* r) {
            return l->val > r->val;
        }
    };

public:
    // SOLUTION
    ListNode* mergeKLists(vector<ListNode*> &lists) {
        priority_queue<ListNode*, vector<ListNode*>, compare> heap;
        ListNode head;
        ListNode *c = &head;
        int i, k, n = lists.size();

        for (i=0; i<n; i++)
            if (lists[i]) heap.push(lists[i]);

        while (!heap.empty()) {
            c->next = heap.top();
            heap.pop();
            c = c->next;
            if (c->next) heap.push(c->next);
        }

        return head.next;
    }
};



int main() {
    Solution o;
    ListNode* ll1 = NULL;
    ListNode* ll2 = NULL;
    ListNode* ll3 = NULL;

    // INPUT
    vector<int> li1 = {1,4,5};
    vector<int> li2 = {1,3,4};
    vector<int> li3 = {2,6};
    llInput(&ll1, li1);
    llInput(&ll2, li2);
    llInput(&ll3, li3);
    vector<ListNode*> lists = {ll1, ll2, ll3};

    // OUTPUT
    auto result = o.mergeKLists(lists);
    llOutput(result);

    return 0;
}
