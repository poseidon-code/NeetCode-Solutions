// 153: Find Minimum in Rotated Sorted Array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/


#include <iostream>
#include <vector>

using namespace std;


class Solution {
public:
    // SOLUTION
    int findMin(vector<int>& nums) {
        int l = 0;
        int r = nums.size() - 1;
        
        while (l<r && nums[l]>nums[r]) {
            int m = l + (r-l) / 2;

            if (nums[m] > nums[r])
                l = m + 1;
            else
                r = m;
        }

        return nums[l];
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> nums = {3,4,5,1,2};

    // OUTPUT
    auto result = o.findMin(nums);
    cout << result << endl;

    return 0;
}