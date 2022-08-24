// 22: Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/


import java.util.ArrayList;
import java.util.List;


class Solution {
    // SOLUTION
    static List<String> result = new ArrayList<>();
    public String[] generateParentheses(int n) {
        generate(n, 0, 0, "");
        return result.toArray(new String[0]);
    }

    private static void generate(int n, int open, int close, String s) {
        if (open==n && close==n) {
            result.add(s);
            return;
        }
        if (open < n)
            generate(n, open+1, close, s+'(');
        if (open > close)
            generate(n, open, close+1, s+')');
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int n = 3;

        // OUTPUT
        var result = o.generateParentheses(n);
        System.out.print("["); for (var v : result) System.out.print("\""+v+"\" "); System.out.println("\b]");
    }
}