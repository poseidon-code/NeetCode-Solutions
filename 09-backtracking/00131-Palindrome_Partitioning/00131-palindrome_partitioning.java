// 131: Palindrome Partitioning
// https://leetcode.com/problems/palindrome-partitioning


import java.util.List;
import java.util.ArrayList;

class Solution {
    private void backtrack(String s, int start, List<String> current, List<List<String>> result) {
        if (start == s.length()) {
            result.add(new ArrayList<>(current));
            return;
        }

        for (int i=start; i<s.length(); i++) {
            if (isPalindrome(s, start, i)) {
                String str = s.substring(start, i + 1);
                current.add(str);
                backtrack(s, i + 1, current, result);
                current.remove(current.size() - 1);
            }
        }
    }

    private boolean isPalindrome(String s, int left, int right) {
        while (left < right) {
            if (s.charAt(left) != s.charAt(right)) return false;
            left++;
            right--;
        }

        return true;
    }

    
    // SOLUTION
    public List<List<String>> partition(String s) {
        List<List<String>> result = new ArrayList<>();
        List<String> current = new ArrayList<>();
        backtrack(s, 0, current, result);
        return result;
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        String s = "aab";

        // OUTPUT
        var result = o.partition(s);
        System.out.println(result);
    }
}
