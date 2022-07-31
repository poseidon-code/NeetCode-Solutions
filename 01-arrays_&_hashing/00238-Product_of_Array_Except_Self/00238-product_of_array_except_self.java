// 238 : Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/


import java.util.Arrays;

class Solution {
    // SOLUTION
    public int[] productExceptSelf(int[] nums) {
        int[] result = new int[nums.length];
        Arrays.fill(result, 1);

        for (int i=0, suf=1, pre=1, n=nums.length; i<n; i++) {
            result[i] *= pre;
            pre *= nums[i];
            result[n-1-i] *= suf;
            suf *= nums[n-1-i];
        }

        return result;
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] nums = {1,2,3,4};

        // OUTPUT
        var result = o.productExceptSelf(nums);
        System.out.print("["); for (var v : result) System.out.print(v + ","); System.out.println("\b]");
    }
}