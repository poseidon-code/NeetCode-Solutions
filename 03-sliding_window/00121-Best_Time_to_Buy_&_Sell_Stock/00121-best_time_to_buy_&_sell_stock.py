# 121. Best Time to Buy and Sell Stock
# https://leetcode.com/problems/best-time-to-buy-and-sell-stock/


class Solution:
    def maxProfit(self, prices: list[int]) -> int:
        minValue = prices[0]
        result = 0

        for i in range(len(prices)):
            minValue = min(minValue, prices[i])
            result = max(result, prices[i]-minValue)

        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    prices: list[int] = [7,1,5,3,6,4]

    # OUTPUT
    result = o.maxProfit(prices)
    print(result)
