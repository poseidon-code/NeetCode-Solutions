# 875: Koko Eating Bananas
# https://leetcode.com/problems/koko-eating-bananas/


class Solution:
    # SOLUTION
    def minEatingSpeed(self, piles: list[int], h: int) -> int:
        l, r = 1, max(piles)
        
        while (l < r):
            total = 0
            m = (l + r) // 2

            for p in piles:
                total += (p + m - 1) // m
            if (total > h):
                l = m + 1
            else:
                r = m

        return l



if __name__ == "__main__":
    o = Solution()

    # INPUT
    piles: list[int] = [3,6,7,11]
    h: int = 8

    # OUTPUT
    result = o.minEatingSpeed(piles, h)
    print(result)