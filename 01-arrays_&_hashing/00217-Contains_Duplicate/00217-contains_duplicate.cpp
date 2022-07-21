// 217: Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/

#include<iostream>
#include<vector>
#include<unordered_set>

using namespace std;

class Solution {
public:
    // SOLUTION
    bool containsDuplicate(vector<int>& nums) {
        unordered_set<int> us;
        for (int i=0; i<nums.size(); i++)
            if (us.insert(nums[i]).second == false)
                return true;
        return false;
    }
};

int main() {
    Solution o;

    // INPUT
    vector<int> nums = {1,2,3,1};

    // OUTPUT
    auto result = o.containsDuplicate(nums);
    cout << (result ? "true" : "false") << endl;
    
    return 0;
}
