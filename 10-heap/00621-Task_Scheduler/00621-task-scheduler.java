// 621: Task Scheduler
// https://leetcode.com/problems/task-scheduler


import java.util.PriorityQueue;
import java.util.ArrayList;
import java.util.List;
import java.util.HashMap;
import java.util.Map;


class Solution {
    // SOLUTION
    public int leastInterval(char[] tasks, int n) {
        Map<Character, Integer> counts = new HashMap<Character, Integer>();
        for (char t : tasks) counts.put(t, counts.getOrDefault(t, 0) + 1);

        PriorityQueue<Integer> pq = new PriorityQueue<Integer>((a, b) -> b - a);
        pq.addAll(counts.values());

        int result = 0;
        int cycle = n + 1;

        while (!pq.isEmpty()) {
            int time = 0;
            List<Integer> tmp = new ArrayList<Integer>();

            for (int i = 0; i < cycle; i++) {
                if (!pq.isEmpty()) {
                    tmp.add(pq.poll());
                    time++;
                }
            }

            for (int cnt : tmp)
                if (--cnt > 0)
                    pq.offer(cnt);

            result += !pq.isEmpty() ? cycle : time;
        }

        return result;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        char[] tasks = {'A','C','A','B','D','B'};
        int n = 1;

        // OUTPUT
        var result = o.leastInterval(tasks, n);
        System.out.println(result);
    }
}