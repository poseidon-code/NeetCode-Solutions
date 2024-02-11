// 39: Combination Sum
// https://leetcode.com/problems/combination-sum/



#include <iostream>
#include <vector>

using namespace std;

class Solution {
private:
    void backtrack(vector<int>& candidates, int target, int start, vector<int>& current, vector<vector<int>>& result) {
        if (start >= candidates.size() || target < 0) return;

        if (target == 0) {
            result.push_back(current);
            return;
        }

        current.push_back(candidates[start]);
        backtrack(candidates, target - candidates[start], start, current, result);
        current.pop_back();
        backtrack(candidates, target, start+1, current, result);
    }

public:
    // SOLUTOIN
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        vector<int> current;
        vector<vector<int>> result;
        backtrack(candidates, target, 0, current, result);
        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> candidates = {2,3,6,7};
    int target = 7;

    // OUTPUT
    auto result = o.combinationSum(candidates, target);

    cout << "["; for (auto x : result) {
        cout << "["; for (auto y : x) {
            cout << y << ",";
        } cout << "\b],";
    } cout << "\b]" << endl;

    return 0;
}
