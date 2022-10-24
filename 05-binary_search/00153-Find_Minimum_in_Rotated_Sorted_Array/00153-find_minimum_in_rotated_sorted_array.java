// 153: Find Minimum in Rotated Sorted Array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/


class Solution {
    // SOLUTION
    public int findMin(int[] nums) {
        int l = 0;
        int r = nums.length - 1;
        
        while (l<r && nums[l]>nums[r]) {
            int m = l + (r-l) / 2;

            if (nums[m] > nums[r])
                l = m + 1;
            else
                r = m;
        }

        return nums[l];
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        
        // INPUT
        int[] nums = {3,4,5,1,2};

        // OUTPUT
        var result = o.findMin(nums);
        System.out.println(result);
    }
}