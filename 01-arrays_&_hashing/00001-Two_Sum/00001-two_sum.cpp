// 1: Two Sum
// https://leetcode.com/problems/two-sum/

#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    // SOLUTION
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> map;

        for (int i=0;;i++) {
            auto it = map.find(target - nums[i]);

            if (it != map.end())
                return {it->second, i};
            
            map[nums[i]] = i;
        }
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> nums = {3,2,4};
    int target = 6;

    // OUTPUT
    auto result = o.twoSum(nums, target);
    cout<<"["<<result[0]<<", "<<result[1]<<"]"<<endl;

    return 0;
}
