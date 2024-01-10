// 39: Combination Sum
// https://leetcode.com/problems/combination-sum/



import java.util.List;
import java.util.ArrayList;

class Solution {
    private void backtrack(int[] candidates, int target, int start, List<Integer> current, List<List<Integer>> result) {
        if (start >= candidates.length || target < 0) return;

        if (target == 0) {
            result.add(new ArrayList<Integer>(current));
            return;
        }

        current.add(candidates[start]);
        backtrack(candidates, target - candidates[start], start, current, result);
        current.remove(current.size()-1);
        backtrack(candidates, target, start+1, current, result);
    }

    // SOLUTION
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> result = new ArrayList<>();
        List<Integer> current = new ArrayList<>();
        backtrack(candidates, target, 0, current, result);
        return result;
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] candidates = {2,3,6,7};
        int target = 7;

        // OUTPUT
        var result = o.combinationSum(candidates, target);
        System.out.println(result);
    }
}
