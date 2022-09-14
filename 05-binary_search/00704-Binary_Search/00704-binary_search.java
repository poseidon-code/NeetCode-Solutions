// 704: Binary Search
// https://leetcode.com/problems/binary-search/


class Solution {
    // SOLUTION
    int search(int[] nums, int target) {
        int l = 0;
        int h = nums.length - 1;

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


    public static void main(String[] args) {
        Solution o = new Solution();
        
        // INPUT
        int[] nums = {-1,0,3,5,9,12};
        int target = 9;

        // OUTPUT
        var result = o.search(nums, target);
        System.out.println(result);
    }
}