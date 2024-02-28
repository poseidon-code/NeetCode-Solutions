// 17: Letter Combinations Of A Phone Number
// https://leetcode.com/problems/letter-combinations-of-a-phone-number


import java.util.List;
import java.util.Map;
import java.util.ArrayList;


class Solution {
    private void backtrack(String digits, int index,  Map<Character, String> m, String current, List<String> result) {
        if (index == digits.length()) {
            result.add(current);
            return;
        }

        String s = m.get(digits.charAt(index));
        for (int i=0; i<s.length(); i++) {
            current += (s.charAt(i));
            backtrack(digits, index + 1, m, current, result);
            current = current.substring(0, current.length() - 1);
        }
    }

    
    // SOLUTION
    public List<String> letterCombinations(String digits) {
        if (digits.length() == 0) {
            return new ArrayList<String>();
        }

        Map<Character, String> m = Map.of(
            '2', "abc",
            '3', "def",
            '4', "ghi",
            '5', "jkl",
            '6', "mno",
            '7', "pqrs",
            '8', "tuv",
            '9', "wxyz"
        );

        List<String> result = new ArrayList<>();
        String current = "";
        backtrack(digits, 0, m, current, result);
        return result;
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        String digits = "23";

        // OUTPUT
        var result = o.letterCombinations(digits);
        System.out.println(result);
    }
}
