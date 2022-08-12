// 11: Container with Most Water
// https://leetcode.com/problems/container-with-most-water/


class Solution {
    // SOLUTION
    public int maxArea(int[] height) {
        int l = 0;
        int r = height.length - 1;
        int result = 0;

        while (l<r) {
            result = Math.max(result, (r-l)*Math.min(height[l], height[r]));
            if (height[l]<height[r]) l++;
            else r--;
        }

        return result;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] height = {1,8,6,2,5,4,8,3,7};

        // OUTPUT
        var result = o.maxArea(height);
        System.out.println(result);
    }
}