// 424: Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/


class Solution {
    // SOLUTION
    int characterReplacement(String s, int k) {
        int[] count = new int[26];
        int maxCount = 0;
        int i = 0;
        int j = 0;

        int result = 0;

        while (j<s.length()) {
            count[s.charAt(j)-'A']++;
            maxCount = Math.max(maxCount, count[s.charAt(j)-'A']);

            if (j-i+1-maxCount > k) {
                count[s.charAt(i)-'A']--;
                i++;
            }

            result = Math.max(result, j-i+1);
            j++;
        }

        return result;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        String s = "ABAB";
        int k = 2;

        // OUTPUT
        var result = o.characterReplacement(s, k);
        System.out.println(result);
    }
}