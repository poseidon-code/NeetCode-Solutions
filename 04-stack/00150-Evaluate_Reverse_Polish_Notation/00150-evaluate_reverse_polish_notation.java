// 150: Evaluate Reverse Polish Notation
// https://leetcode.com/problems/evaluate-reverse-polish-notation/


import java.util.Stack;


class Solution {
    // SOLUTION
    public int evalRPN(String[] tokens) {
        Stack<Integer> result = new Stack<>();

        for (var s : tokens) {
            if(s.length()>1 || Character.isDigit(s.charAt(0))) {
                result.push(Integer.parseInt(s));
            } else {
                var x2 = result.peek(); result.pop();
                var x1 = result.peek(); result.pop();
                
                switch(s.charAt(0)) {
                    case '+': x1+=x2; break;
                    case '-': x1-=x2; break;
                    case '*': x1*=x2; break;
                    case '/': x1/=x2; break;
                }

                result.push(x1);
            }
        }

        return result.peek();
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        String[] tokens = {"2","1","+","3","*"};

        // OUTPUT
        var result = o.evalRPN(tokens);
        System.out.println(result);
    }
}