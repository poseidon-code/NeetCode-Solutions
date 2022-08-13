// 42: Trapping Rain Water
// https://leetcode.com/problems/trapping-rain-water/


class Solution {
    // SOLUTION
    int trap(int[] height) {
        int l = 0;
        int r = height.length - 1;

        int maxLeft = height[l];
        int maxRight = height[r];

        int result = 0;

        while (l<r) {
            if (maxLeft<=maxRight) {
                l++;
                maxLeft = Math.max(maxLeft, height[l]);
                result += maxLeft - height[l];
            } else {
                r--;
                maxRight = Math.max(maxRight, height[r]);
                result += maxRight - height[r];
            }
        }
        
        return result;
    }



    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] height = {4,2,0,3,2,5};

        // OUTPUT
        var result = o.trap(height);
        System.out.println(result);
    }
}