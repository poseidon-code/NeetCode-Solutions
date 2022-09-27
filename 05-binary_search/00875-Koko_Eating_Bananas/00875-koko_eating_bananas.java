// 875: Koko Eating Bananas
// https://leetcode.com/problems/koko-eating-bananas/

import java.util.Arrays;

class Solution {
    // SOLUTION
    public int minEatingSpeed(int[] piles, int h) {
        int l = 1, r = Arrays.stream(piles).max().getAsInt();
        
        while (l < r) {
            int total = 0;
            int m = (l + r) / 2;

            for (int p : piles)
                total += (p + m - 1) / m;
            if (total > h)
                l = m + 1;
            else
                r = m;
        }

        return l;
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        
        // INPUT
        int[] piles = {3,6,7,11};
        int h = 8;

        // OUTPUT
        var result = o.minEatingSpeed(piles, h);
        System.out.println(result);
    }
}