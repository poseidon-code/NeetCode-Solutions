// 167: Two Sum II
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    // SOLUTION
    vector<int> twoSum(vector<int>& numbers, int target) {
        int l = 0;
        int r = numbers.size() - 1;
        
        while (l<r) {
            int s = numbers[l] + numbers[r];
            if (s==target) return {l+1, r+1};
            else if (s<target) l++;
            else r--;
        }

        return {};
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> numbers = {2,7,11,15};
    int target = 9;

    // OUTPUT
    auto result = o.twoSum(numbers, target);
    cout<<"["; for (auto v : result) cout<<v<<" "; cout<<"\b]"<<endl;
}