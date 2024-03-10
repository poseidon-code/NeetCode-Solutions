// 703: Kth Largest Element In A Stream
// https://leetcode.com/problems/kth-largest-element-in-a-stream


import java.util.PriorityQueue;

class KthLargest {
    // SOLUTION
    static PriorityQueue<Integer> pq;
    static int size;

    public KthLargest(int k, int[] nums) {
        size = k;
        pq = new PriorityQueue<>(k);
        for (int n : nums)
            add(n);
    }

    public int add(int val) {
        if (pq.size() < size)
            pq.offer(val);
        else if (pq.peek() < val) {
            pq.poll();
            pq.offer(val);
        }
        return pq.peek();
    }

    
    public static void main(String[] args) {
        // INPUT
        int k = 3;
        int[] nums = {4,5,8,2};
        KthLargest o = new KthLargest(k, nums);

        // OPERATIONS
        System.out.println(o.add(3));
        System.out.println(o.add(5));
        System.out.println(o.add(10));
        System.out.println(o.add(9));
        System.out.println(o.add(4));
    }
}
