// 78: Subsets
// https://leetcode.com/problems/subsets/


import java.util.List;
import java.util.ArrayList;

class Solution {
    private void backtrack(int[] nums, int start, List<Integer> current, List<List<Integer>> result) {
        result.add(new ArrayList<Integer>(current));

        for (int i = start; i < nums.length; i++) {
            current.add(nums[i]);
            backtrack(nums, i + 1, current, result);
            current.remove(current.size() - 1);
        }
    }

    // SOLUTION
    public List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        List<Integer> current = new ArrayList<>();
        backtrack(nums, 0, current, result);
        return result;
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] nums = {1,2,3};

        // OUTPUT
        var result = o.subsets(nums);

        System.out.print("["); for (var x : result) {
            System.out.print("["); for (var y : x) {
                System.out.print(y + ",");
            } System.out.print("\b],");
        } System.out.println("\b]");
    }
}
