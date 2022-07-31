// 238 : Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/


#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    // SOLUTION
    vector<int> productExceptSelf(vector<int>& nums) {
        vector<int> result(size(nums), 1);

        for (int i=0, suf=1, pre=1, n=size(nums); i<size(nums); i++) {
            result[i] *= pre;
            pre *= nums[i];
            result[n-1-i] *= suf;
            suf *= nums[n-1-i];
        }

        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> nums = {1,2,3,4};

    // OUTPUT
    auto result = o.productExceptSelf(nums);
    cout << "["; for (auto v : result) cout << v << ","; cout << "\b]" << endl;

    return 0;
}