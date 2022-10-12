// 33. Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/


#include <iostream>
#include <vector>

using namespace std;


class Solution {
public:
    // SOLUTION
    int search(vector<int>& nums, int target) {
        int l = 0;
        int h = nums.size();

        while (l < h) {
            int m = (l+h) / 2;
            if (target < nums[0] < nums[m])
                l = m+1;
            else if (target >= nums[0] > nums[m])
                h = m;
            else if (nums[m] < target)
                l = m+1;
            else if (nums[m] > target)
                h = m;
            else
                return m;
        }

        return -1;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> nums = {4,5,6,7,0,1,2};
    int target = 0;

    // OUTPUT
    auto result = o.search(nums, target);
    cout << result << endl;

    return 0;
}