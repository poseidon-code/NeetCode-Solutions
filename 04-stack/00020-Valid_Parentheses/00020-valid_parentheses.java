// 20: Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/


import java.util.Stack;

class Solution {
    // SOLUTION
    public boolean isValid (String s) {
        Stack<Character> parentheses = new Stack<>();

        for (var c : s.toCharArray()) {
            switch(c) {
                case '{': parentheses.push('}'); break;
                case '[': parentheses.push(']'); break;
                case '(': parentheses.push(')'); break;
                default:
                    if (parentheses.isEmpty() || c!=parentheses.pop())
                        return false;
            }
        }

        return parentheses.isEmpty();
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        String s = "()[]{}";

        // OUTPUT
        var result = o.isValid(s);
        System.out.println(result);
    }
}