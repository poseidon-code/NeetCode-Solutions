// 40: Combination Sum Ii
// https://leetcode.com/problems/combination-sum-ii


#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
private:
    void backtrack(vector<int>& candidates, int target, int sum, int start, vector<int>& current, vector<vector<int>>& result) {
        if (sum > target) return;

        if (sum == target) {
            result.push_back(current);
            return;
        }

        for (int i=start; i<candidates.size(); i++) {
            if (i>start && candidates[i] == candidates[i-1]) continue;
            current.push_back(candidates[i]);
            backtrack(candidates, target, sum + candidates[i], i+1, current, result);
            current.pop_back();
        }
    }

public:
    // SOLUTOIN
    vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
        std::sort(candidates.begin(), candidates.end());
        vector<int> current;
        vector<vector<int>> result;
        backtrack(candidates, target, 0, 0, current, result);
        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> candidates = {2,5,2,1,2};
    int target = 5;

    // OUTPUT
    auto result = o.combinationSum2(candidates, target);

    cout << "["; for (auto x : result) {
        cout << "["; for (auto y : x) {
            cout << y << ",";
        } cout << "\b],";
    } cout << "\b]" << endl;

    return 0;
}
