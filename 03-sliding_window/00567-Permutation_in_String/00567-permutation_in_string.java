// 567: Permutation in String
// https://leetcode.com/problems/permutation-in-string/


import java.util.Arrays;


class Solution {
    // SOLUTION
    boolean checkInclusion(String s1, String s2) {
        if (s1.length() > s2.length()) return false;
        
        int[] s1Count = new int[26]; for (var c : s1.toCharArray()) s1Count[c-'a']++;
        int[] s2Count = new int[26];
        int i = 0;
        int j = 0;
        
        while (j < s2.length()) {
            s2Count[s2.charAt(j)-'a']++;
            
            if (j-i+1 == s1.length()) 
                if (Arrays.equals(s1Count, s2Count)) 
                    return true;
            if (j-i+1 < s1.length()) 
                j++;
            else {
                s2Count[s2.charAt(i)-'a']--;
                i++;
                j++;
            }
        }

        return false;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        String s1 = "ab";
        String s2 = "eidbaooo";

        // OUTPUT
        var result = o.checkInclusion(s1, s2);
        System.out.println(result);
    }
}