// 739: Daily Temperstures
// https://leetcode.com/problems/daily-temperatures/


import java.util.Stack;

class Solution {
    // SOLUTION
    public int[] dailyTemperatures(int[] temperatures) {
        Stack<Integer> s = new Stack<>();
        int[] result = new int[temperatures.length];

        for (int i=0; i<temperatures.length; i++) {
            while (!s.isEmpty() && temperatures[i]>temperatures[s.peek()]) {
                result[s.peek()] = i - s.peek();
                s.pop();
            }
            s.push(i);
        }

        return result;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] temperatures = {73,74,75,71,69,72,76,73};

        // OUTPUT
        var result = o.dailyTemperatures(temperatures);
        System.out.print("["); for (var v : result) System.out.print(v+" "); System.out.println("\b]");
    }
}