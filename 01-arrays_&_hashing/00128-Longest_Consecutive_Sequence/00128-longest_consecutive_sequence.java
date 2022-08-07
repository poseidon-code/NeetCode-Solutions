// 00128: Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/

import java.util.HashMap;

class Solution {
    // SOLUTION
    public int longestConsecutive(int nums[]) {
        int result = 0;
        HashMap<Integer, Integer> tmap = new HashMap<Integer, Integer>();
        
        for (int i : nums) {
            if (tmap.containsKey(i)) continue;

            int left = tmap.getOrDefault(i-1, 0);
            int right = tmap.getOrDefault(i+1, 0);
            int sum = left + right + 1;
            result = Math.max(result, sum);

            if (left>0) tmap.put(i-left, sum);
            if (right>0) tmap.put(i+right, sum);
            tmap.put(i, sum);
        }
        
        return result;
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int nums[] = {100,4,200,1,3,2};

        // OUTPUT
        var result = o.longestConsecutive(nums);
        System.out.println(result);
    }
}