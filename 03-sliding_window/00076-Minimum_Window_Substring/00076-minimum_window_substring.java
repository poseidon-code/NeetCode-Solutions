// 76: Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/


import java.util.HashMap;


class Solution {
    // SOLUTION
    public String minWindow(String s, String t) {
        HashMap<Character, Integer> m = new HashMap<>();
        for (char c : t.toCharArray()) m.put(c, m.getOrDefault(c, 0) + 1);

        int i = 0;
        int j = 0;
        int counter = t.length();
        
        int minStart = 0;
        int minLength = s.length() + 1;
        
        while (j < s.length()) {
            if (m.containsKey(s.charAt(j))) {
                m.put(s.charAt(j), m.get(s.charAt(j))-1);
                counter--;
            }
            j++;

            while (counter==0) {
                if (j-i < minLength) {
                    minStart = i;
                    minLength = j-i;
                }
                if (m.containsKey(s.charAt(i))) {
                    m.put(s.charAt(i), m.get(s.charAt(i))+1);
                    counter++;
                }
                i++;
            }
        }
        
        if (minLength != s.length())
            return s.substring(minStart, minStart+minLength);
        return "";
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        String s = "ADOBECODEBANC";
        String t = "ABC";

        // OUTPUT
        var result = o.minWindow(s, t);
        System.out.println(result);
    }
}