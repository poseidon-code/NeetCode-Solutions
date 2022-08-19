// 239: Sliding Window Maximum
// https://leetcode.com/problems/sliding-window-maximum/


import java.util.Deque;
import java.util.LinkedList;


class Solution {
    // SOLUTION
    public int[] maxSlidingWindow(int[] nums, int k) {
        Deque<Integer> dq = new LinkedList<>();
        int[] result = new int[nums.length-k+1];

        int j=0;

        for (int i=0; i<nums.length; i++) {
            if (!dq.isEmpty() && dq.peekFirst() < i-k+1) dq.pollFirst();
            
            while (!dq.isEmpty() && nums[i]>nums[dq.peekLast()]) 
                dq.pollLast();
            
            dq.offer(i);
            
            if (i >= k-1) 
                result[j++] = nums[dq.peekFirst()];
        }

        return result;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int nums[] = {1,3,-1,-3,5,3,6,7};
        int k = 3;

        // OUTPUT
        var result = o.maxSlidingWindow(nums, k);
        System.out.print("["); for (var v : result) System.out.print(v+" "); System.out.println("\b]");
    }
}