// 287: Find the Duplicate Number
// https://leetcode.com/problems/find-the-duplicate-number/


#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    // SOLUTION
    int findDuplicate(vector<int>& nums) {        
        int slow = nums[0], fast = nums[nums[0]];
        
        while (slow != fast) {
            slow = nums[slow];
            fast = nums[nums[fast]];
        }
        
        slow = 0;
        while (slow != fast) {
            slow = nums[slow];
            fast = nums[fast];
        }

        return slow;
    }
};



int main() {
    Solution o;

    // INPUT
    vector<int> nums = {1,3,4,2,2};

    // OUTPUT
    auto result = o.findDuplicate(nums);
    cout << result << endl;

    return 0;
}
