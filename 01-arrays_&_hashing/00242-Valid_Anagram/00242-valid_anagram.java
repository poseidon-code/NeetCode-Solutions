// 242: Valid Anagram
// https://leetcode.com/problems/valid-anagram/

class Solution {
    // SOLUTION
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) return false;

        int n = s.length();
        int counts[] = new int[26];
        
        for (int i=0; i<n; i++) {
            counts[s.charAt(i) - 'a']++;
            counts[t.charAt(i) - 'a']--;
        }

        for (int i=0; i<26; i++)
            if (counts[i] != 0) return false;
        
        return true;
    }

    public static void main(String[] args) {
        Solution o = new Solution();
        
        // INPUT
        String s = "rat";
        String t = "cat";

        // OUTPUT
        var result = o.isAnagram(s, t);
        System.out.println((result ? "true" : "false"));
    }
}
