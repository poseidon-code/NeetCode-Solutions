// 00128: Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/

#include <iostream>
#include <vector>
#include <unordered_map>
#include <algorithm>

using namespace std;

class Solution {
public:
    // SOLUTION
    int longestConsecutive(vector<int>& nums) {
        int result = 0;
        unordered_map<int,int> tmap;

        for (int i : nums) {
            if (tmap[i]) continue;
            result = max(
                            result, 
                            tmap[i]
                                = tmap[i + tmap[i+1]]
                                = tmap[i - tmap[i-1]]
                                = tmap[i+1] + tmap[i-1] + 1
                        );
        }

        return result;
    }
};

int main() {
    Solution o;

    // INPUT
    vector<int> nums = {100,4,200,1,3,2};

    // OUTPUT
    auto result = o.longestConsecutive(nums);
    cout<<result<<endl;

    return 0;
}