// 215: Kth Largest Element In An Array
// https://leetcode.com/problems/kth-largest-element-in-an-array


#include <iostream>
#include <queue>
#include <vector>

using namespace std;


class Solution {
public:
    // SOLUTION
    int findKthLargest(vector<int>& nums, int k) {
        priority_queue<int, vector<int>, greater<int>> pq;
        
        for (int num : nums) {
            pq.push(num);
            if (pq.size() > k) {
                pq.pop();
            }
        }
        
        return pq.top();
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> nums = {3,2,1,5,6,4};
    int k = 2;

    // OUTPUT
    auto result = o.findKthLargest(nums, k);
    cout << result << endl;

    return 0;
}
