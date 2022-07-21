// 217: Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/

import java.util.Set;
import java.util.HashSet;

class Solution {
    // SOLUTION
    public boolean containsDuplicate(int[] nums) {
        Set<Integer> set = new HashSet<Integer>();
		
        for(var i : nums)
			if(!set.add(i))
				return true; 
		
        return false;
    }

    public static void main(String[] args) {
        Solution o = new Solution();
        
        // INPUT
        int nums[] = {1,2,3,1};

        // OUTPUT
        var result = o.containsDuplicate(nums);
        System.out.println((result ? "true" : "false"));
    }
}
