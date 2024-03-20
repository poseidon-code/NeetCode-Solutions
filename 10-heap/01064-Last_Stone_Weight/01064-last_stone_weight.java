// 1064: Last Stone Weight
// https://leetcode.com/problems/last-stone-weight

import java.util.PriorityQueue;


class Solution {
    public int lastStoneWeight(int[] stones) {
        PriorityQueue<Integer> pq = new PriorityQueue<>((a, b) -> b - a);
        for (int stone : stones) pq.offer(stone);

        while (pq.size() > 1) {
            int y = pq.poll();
            int x = pq.poll();

            if (x != y) {
                pq.offer(y - x);
            }
        }

        return pq.isEmpty() ? 0 : pq.peek();
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] stones = {2,7,4,1,8,1};

        // OUPUT
        var result = o.lastStoneWeight(stones);
        System.out.println(result);
    }
}