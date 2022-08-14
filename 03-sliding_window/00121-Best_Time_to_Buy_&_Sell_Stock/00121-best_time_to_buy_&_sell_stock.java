// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/


class Solution {
    // SOLUTION
    public int maxProfit(int[] prices) {
        int minValue = prices[0];
        int result = 0;

        for (int i=1; i<prices.length; i++) {
            minValue = Math.min(minValue, prices[i]);
            result = Math.max(result, prices[i] - minValue);
        }

        return result;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] prices = {7,1,5,3,6,4};

        // OUTPUT
        var result = o.maxProfit(prices);
        System.out.println(result);
    }
}