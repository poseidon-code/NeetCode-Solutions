// 215: Kth Largest Element In An Array
// https://leetcode.com/problems/kth-largest-element-in-an-array


import java.util.PriorityQueue;

class Solution {
    // SOLUTION
    public int findKthLargest(int[] nums, int k) {
        PriorityQueue<Integer> pq = new PriorityQueue<Integer>(k + 1);

        for(int num : nums) {
            pq.add(num);
            if (pq.size() > k) {
                pq.poll();
            }
        }

        return pq.poll();
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] nums = {3,2,1,5,6,4};
        int k = 2;

        // OUTPUT
        var result = o.findKthLargest(nums, k);
        System.out.println(result);
    }
}