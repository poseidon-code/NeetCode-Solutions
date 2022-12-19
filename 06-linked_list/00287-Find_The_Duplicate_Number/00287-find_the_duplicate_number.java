// 287: Find the Duplicate Number
// https://leetcode.com/problems/find-the-duplicate-number/


class Solution {
    // SOLUTION
    public int findDuplicate(int[] nums) {
        int slow = nums[0];
        int fast = nums[nums[0]];
        
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


    public static void main(String[] args) {
        Solution o = new Solution();
        
        // INPUT
        int[] nums = {1,3,4,2,2};

        // OUTPUT
        var result = o.findDuplicate(nums);
        System.out.println(result);
    }
}