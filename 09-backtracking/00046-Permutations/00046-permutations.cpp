// 46: Permutations
// https://leetcode.com/problems/permutations/



#include <iostream>
#include <vector>

using namespace std;

class Solution {
private:
    void backtrack(vector<int>& nums, int start, vector<vector<int>>& result) {
        if (start == nums.size()) {
            result.push_back(nums);
            return;
        }

        for (int i=start; i<nums.size(); i++) {
            swap(nums[i], nums[start]);
            backtrack(nums, start + 1, result);
            swap(nums[i], nums[start]);
        }
    }

public:
    // SOLUTOIN
    vector<vector<int>> permute(vector<int>& nums) {
        vector<vector<int>> result;
        backtrack(nums, 0, result);
        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> nums = {0,1};

    // OUTPUT
    auto result = o.permute(nums);

    cout << "["; for (auto x : result) {
        cout << "["; for (auto y : x) {
            cout << y << ",";
        } cout << "\b],";
    } cout << "\b]" << endl;

    return 0;
}
