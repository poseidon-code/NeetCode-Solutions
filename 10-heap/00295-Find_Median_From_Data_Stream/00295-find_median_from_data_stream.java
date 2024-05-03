// 295: Find Median From Data Stream
// https://leetcode.com/problems/find-median-from-data-stream

import java.util.Queue;
import java.util.PriorityQueue;


// SOLUTION
class MedianFinder {
    private Queue<Integer> lower;
    private Queue<Integer> higher;

    public MedianFinder() {
        lower = new PriorityQueue<>((a, b) -> b - a);
        higher = new PriorityQueue<>((a, b) -> a - b);
    }

    public void addNum(int num) {
        lower.add(num);
        if (
            lower.size() - higher.size() > 1 ||
            !higher.isEmpty() &&
            lower.peek() > higher.peek()
        ) {
            higher.add(lower.poll());
        }

        if (higher.size() - lower.size() > 1) {
            lower.add(higher.poll());
        }
    }

    public double findMedian() {
        if (lower.size() == higher.size()) {
            return (double) (higher.peek() + lower.peek()) / 2;
        } else if (lower.size() > higher.size()) {
            return (double) lower.peek();
        } else {
            return (double) higher.peek();
        }
    }

    public static void main(String[] args) {
        MedianFinder o = new MedianFinder();

        // OPERATIONS
        o.addNum(1);
        o.addNum(2);
        System.out.println(o.findMedian());
        o.addNum(3);
        System.out.println(o.findMedian());
    }
}
