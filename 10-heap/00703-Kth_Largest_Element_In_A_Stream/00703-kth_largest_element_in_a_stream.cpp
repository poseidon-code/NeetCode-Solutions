// 703: Kth Largest Element In A Stream
// https://leetcode.com/problems/kth-largest-element-in-a-stream


#include <iostream>
#include <queue>
#include <vector>

using namespace std;


class KthLargest {
// SOLUTION
private:
    priority_queue<int, vector<int>, greater<int>> pq;
    int size;

public:
    KthLargest(int k, vector<int>& nums) {
        this->size = k;
        for (auto n : nums) {
            pq.push(n);
            if (pq.size() > k) pq.pop();
        }
    }

    int add(int val) {
        pq.push(val);
        if (pq.size() > size) pq.pop();
        return pq.top();
    }
};


int main() {
    // INPUT
    int k = 3;
    vector<int> nums = {4,5,8,2};
    KthLargest o(k, nums);


    // OPERATIONS
    cout << o.add(3) << endl;
    cout << o.add(5) << endl;
    cout << o.add(10) << endl;
    cout << o.add(9) << endl;
    cout << o.add(4) << endl;

    return 0;
}
