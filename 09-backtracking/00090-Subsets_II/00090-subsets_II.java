// 90: Subsets Ii
// https://leetcode.com/problems/subsets-ii


import java.util.List;
import java.util.ArrayList;

class Solution {
    private void backtrack(int[] nums, int start, List<Integer> current, List<List<Integer>> result) {
        result.add(new ArrayList<Integer>(current));

        for (int i = start; i < nums.length; i++) {
            if (i>start && nums[i] == nums[i-1]) continue;
            current.add(nums[i]);
            backtrack(nums, i + 1, current, result);
            current.remove(current.size() - 1);
        }
    }

    // SOLUTION
    public List<List<Integer>> subsetsWithDup(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        List<Integer> current = new ArrayList<>();
        backtrack(nums, 0, current, result);
        return result;
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] nums = {1,2,2};

        // OUTPUT
        var result = o.subsetsWithDup(nums);
        System.out.println(result);
    }
}
