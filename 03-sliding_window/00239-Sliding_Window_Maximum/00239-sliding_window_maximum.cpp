// 239: Sliding Window Maximum
// https://leetcode.com/problems/sliding-window-maximum/


#include <iostream>
#include <vector>
#include <deque>

using namespace std;


class Solution {
public:
    // SOLUTION
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        deque<int> dq;
        vector<int> result;

        int i=0, j=0;

        while (j < nums.size()) {
            while (!dq.empty() && nums[dq.back()]<nums[j])
                dq.pop_back();
            dq.push_back(j);
            
            if (i > dq.front()) dq.pop_front();
            
            if (j+1 >= k) {
                result.push_back(nums[dq.front()]);
                i++;
            }
            j++;
        }

        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> nums = {1,3,-1,-3,5,3,6,7};
    int k = 3;

    // OUTPUT
    auto result = o.maxSlidingWindow(nums, k);
    cout<<"["; for (auto v : result) cout<<v<<" "; cout<<"\b]"<<endl;

    return 0;
}