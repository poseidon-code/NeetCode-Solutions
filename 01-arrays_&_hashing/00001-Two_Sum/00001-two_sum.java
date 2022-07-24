// 1: Two Sum
// https://leetcode.com/problems/two-sum/

import java.util.HashMap;
import java.util.Map;

class Solution {
    // SOLUTION
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<Integer, Integer>();

        for (int i=0; i<nums.length; map.put(nums[i], i++))
            if (map.containsKey(target - nums[i]))
                return new int[] {map.get(target - nums[i]), i};
        
        return new int[] {0,0};
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] nums = {3,2,4};
        int target = 6;

        // OUTPUT
        var result = o.twoSum(nums, target);
        System.out.println("[" + result[0] + ", " + result[1] + "]");
    }
}
