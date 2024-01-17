// 46: Permutations
// https://leetcode.com/problems/permutations/



import java.util.List;
import java.util.ArrayList;

class Solution {
    private void backtrack(int[] nums, int start, List<List<Integer>> result) {
        if (start == nums.length) {
            List<Integer> t_list = new ArrayList<Integer>();
            for (var e : nums) t_list.add(e);
            result.add(t_list);
            return;
        }

        for (int i = start; i < nums.length; i++) {
            var t = nums[i];
            nums[i] = nums[start]; nums[start] = t;
            backtrack(nums, start + 1, result);
            t = nums[i];
            nums[i] = nums[start]; nums[start] = t;
        }
    }

    // SOLUTION
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        backtrack(nums, 0, result);
        return result;
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] nums = {0,1};

        // OUTPUT
        var result = o.permute(nums);
        System.out.println(result);
    }
}
