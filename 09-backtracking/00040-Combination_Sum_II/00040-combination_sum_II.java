// 40: Combination Sum Ii
// https://leetcode.com/problems/combination-sum-ii


import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

class Solution {
    private void backtrack(int[] candidates, int target, int sum, int start, List<Integer> current, List<List<Integer>> result) {
        if (sum > target) return;

        if (sum == target) {
            result.add(new ArrayList<Integer>(current));
            return;
        }

        for (int i=start; i<candidates.length; i++) {
            if (i>start && candidates[i] == candidates[i-1]) continue;
            current.add(candidates[i]);
            backtrack(candidates, target, sum + candidates[i], i+1, current, result);
            current.remove(current.size()-1);
        }
    }

    // SOLUTION
    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        Arrays.sort(candidates);
        List<List<Integer>> result = new ArrayList<>();
        List<Integer> current = new ArrayList<>();
        backtrack(candidates, target, 0, 0, current, result);
        return result;
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] candidates = {2,5,2,1,2};
        int target = 5;

        // OUTPUT
        var result = o.combinationSum2(candidates, target);
        System.out.println(result);
    }
}
