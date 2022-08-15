// 3: Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

import java.util.HashSet;

class Solution {
    // SOLUTION
    public int lengthOfLongestSubstring(String s) {
        HashSet<Character> letters = new HashSet<>();
        int i=0, j=0;
        int result = 0;

        for (; j<s.length(); j++) {
            while (letters.contains(s.charAt(j))) {
                letters.remove(s.charAt(i));
                i++;
            } 
            letters.add(s.charAt(j));
            result = Math.max(result, j-i+1);
        }

        return result;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        String s = "abcabcbb";

        // OUTPUT
        var result = o.lengthOfLongestSubstring(s);
        System.out.println(result);
    }
}