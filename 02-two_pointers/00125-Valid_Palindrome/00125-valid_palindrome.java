// 125: Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/

class Solution {
    // SOLUTON
    public boolean isPalindrome(String s) {
        int l = 0;
        int r = s.length() - 1;

        while (l < r) {
            while (!Character.isLetterOrDigit(s.charAt(l)) && l<r) l++;
            while (!Character.isLetterOrDigit(s.charAt(r)) && l<r) r--;
            if (Character.toLowerCase(s.charAt(l)) != Character.toLowerCase(s.charAt(r))) return false;
            
            l++;
            r--;
        }

        return true;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        String s = "race a car";

        // OUTPUT
        var result = o.isPalindrome(s);
        System.out.println(result);        
    }
}