// 15: 3Sum
// https://leetcode.com/problems/3sum/


import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;


class Solution {
    // SOLUTION
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> result = new ArrayList<>();

        for (int i=0; i<nums.length; i++) {
            if (i>0 && nums[i]==nums[i-1]) continue;

            int l = i+1;
            int r = nums.length - 1;

            while (l<r) {
                int s = nums[i] + nums[l] + nums[r];
                if (s>0) r--;
                else if (s<0) l++;
                else {
                    result.add(Arrays.asList(nums[i], nums[l], nums[r]));
                    while (nums[l]==nums[l+1]) l++;
                    while (nums[r]==nums[r-1]) r--;
                    l++;
                    r--;
                }
            }
        }

        return result;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] nums = {-1,0,1,2,-1,-4};

        // OUTPUT
        var result = o.threeSum(nums);
        
        System.out.print("[");
        for (var x : result) {
            System.out.print("[");
            for (var y : x) {
                System.out.print(y+" ");
            }
            System.out.print("\b] ");
        }
        System.out.println("\b]");
    }
}