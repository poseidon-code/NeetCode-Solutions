// 973: K Closest Points To Origin
// https://leetcode.com/problems/k-closest-points-to-origin


import java.util.PriorityQueue;


class Solution {
    // SOLUTION
    public int[][] kClosest(int[][] points, int k) {
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) ->
            Integer.compare(
                (a[0] * a[0] + a[1] * a[1]),
                (b[0] * b[0] + b[1] * b[1])
            )
        );

        for (int[] point : points) {
            pq.add(point);
        }

        int[][] result = new int[k][2];
        for (int i = 0; i < k; i++) {
            int[] e = pq.poll();
            result[i][0] = e[0];
            result[i][1] = e[1];
        }

        return result;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[][] points = {{1,3},{-2,2}};
        int k = 1;

        // OUTPUT
        var result = o.kClosest(points, k);

        System.out.print("["); for (var x : result) {
            System.out.print("["); for (var y : x) {
                System.out.print(y + ",");
            } System.out.print("\b],");
        } System.out.println("\b]");
    }
}
