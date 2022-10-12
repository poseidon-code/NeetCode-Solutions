// 33. Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/


class Solution {
    // SOLUTION
    public int search(int[] nums, int target) {
        int l = 0;
        int h = nums.length;

        while (l < h) {
            int m = (l+h) / 2;
            if (target < nums[0] && nums[0] < nums[m])
                l = m+1;
            else if (target >= nums[0] && nums[0] > nums[m])
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


    public static void main(String[] args) {
        Solution o = new Solution();
        
        // INPUT
        int[] nums = {4,5,6,7,0,1,2};
        int target = 0;

        // OUTPUT
        var result = o.search(nums, target);
        System.out.println(result);
    }
}