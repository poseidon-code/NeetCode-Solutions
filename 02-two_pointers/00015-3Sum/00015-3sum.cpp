// 15: 3Sum
// https://leetcode.com/problems/3sum/

#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    // SOLUTION
    vector<vector<int>> threeSum(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        vector<vector<int>> result;

        for (int i=0; i<nums.size(); i++) {
            if (i>0 && nums[i]==nums[i-1]) continue;

            int l = i+1;
            int r = nums.size() - 1;

            while (l<r) {
                int s = nums[i] + nums[l] + nums[r];
                if (s>0) r--;
                else if (s<0) l++;
                else {
                    result.push_back({nums[i], nums[l], nums[r]});
                    while (nums[l]==nums[l+1]) l++;
                    while (nums[r]==nums[r-1]) r--;
                    l++;
                    r--;
                }
            }
        }

        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> nums = {-1,0,1,2,-1,-4};

    // OUTPUT
    auto result = o.threeSum(nums);

    cout<<"["; for (auto x : result) {
        cout<<"["; for (auto y : x) cout<<y<<" "; cout<<"\b] ";
    } cout<<"\b]"<<endl;
    
    return 0;
}