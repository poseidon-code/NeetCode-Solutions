// 84: Largest Rectangle in Histogram
// https://leetcode.com/problems/largest-rectangle-in-histogram/


import java.util.Stack;


class Solution {
    // SOLUTION
    public int largestRectangleArea(int[] heights) {
        Stack<Integer> stack = new Stack<>();
        int l = heights.length;
        int result = 0;

        for (int i=0; i<=l; i++) {
            int h = i==l ? 0 : heights[i];

            if (stack.isEmpty() || h>=heights[stack.peek()])
                stack.push(i);
            else {
                int tp = stack.pop();
                result = Math.max(result, heights[tp]*(stack.isEmpty() ? i : i-1-stack.peek()));
                i--;

            }
        }

        return result;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] heights = {2,1,5,6,2,3};

        // OUTPUT
        var result = o.largestRectangleArea(heights);
        System.out.println(result);
    }
}