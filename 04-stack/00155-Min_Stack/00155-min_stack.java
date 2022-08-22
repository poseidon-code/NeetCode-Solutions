// 155: Min stack
// https://leetcode.com/problems/min-stack/


import java.util.Stack;


class Solution {
    // SOLUTION
    static class MinStack {
        Stack<Integer> s1 = new Stack<>();
        Stack<Integer> s2 = new Stack<>();

        public void push(int x) {
            s1.push(x);
            if (s2.isEmpty() || x <= getMin())
                s2.push(x);
        }

        public void pop() {
            if (s1.peek() == getMin())
                s2.pop();
            s1.pop();
        }

        public int top() {
            return s1.peek();
        }

        public int getMin() {
            return s2.peek();
        }
    }


    public static void main(String[] args) {
        MinStack o = new MinStack();

        // OPERATIONS
        o.push(-2);
        o.push(0);
        o.push(-3);
        System.out.print(o.getMin() + " ");
        o.pop();
        o.top();
        System.out.print(o.getMin() + " ");

        System.out.println();
    }
}