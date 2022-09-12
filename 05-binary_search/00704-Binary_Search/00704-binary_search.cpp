// 704: Binary Search
// https://leetcode.com/problems/binary-search/


#include <iostream>
#include <vector>

using namespace std;


class Solution {
public:
    // SOLUTION
    int search(vector<int>& nums, int target) {
        int l = 0;
        int h = nums.size() - 1;

        while (l <= h) {
            int m = l + (h-l)/2;
            
            if (nums[m] < target)
                l = m+1;
            else if (nums[m] > target)
                h = m-1;
            else
                return m;
        }

        return -1;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> nums = {-1,0,3,5,9,12};
    int target = 9;

    // OUTPUT
    auto result = o.search(nums, target);
    cout<<result<<endl;

    return 0;
}