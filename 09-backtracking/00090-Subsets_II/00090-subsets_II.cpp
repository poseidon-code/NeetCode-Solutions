// 90: Subsets Ii
// https://leetcode.com/problems/subsets-ii



#include <iostream>
#include <vector>

using namespace std;

class Solution {
private:
    void backtrack(vector<int>& nums, int start, vector<int>& current, vector<vector<int>>& result) {
        result.push_back(current);

        for (int i=start; i<nums.size(); i++) {
            if (i>start && nums[i] == nums[i-1]) continue;
            current.push_back(nums[i]);
            backtrack(nums, i + 1, current, result);
            current.pop_back();
        }
    }

public:
    // SOLUTOIN
    vector<vector<int>> subsetsWithDup(vector<int>& nums) {
        vector<int> current;
        vector<vector<int>> result;
        backtrack(nums, 0, current, result);
        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> nums = {1,2,2};

    // OUTPUT
    auto result = o.subsetsWithDup(nums);

    cout << "["; for (auto x : result) {
        cout << "["; for (auto y : x) {
            cout << y << ",";
        } cout << "\b],";
    } cout << "\b]" << endl;

    return 0;
}
