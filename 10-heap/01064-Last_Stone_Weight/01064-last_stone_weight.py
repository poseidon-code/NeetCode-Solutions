# 1064: Last Stone Weight
# https://leetcode.com/problems/last-stone-weight


import heapq

class Solution:
    def lastStoneWeight(self, stones: list[int]) -> int:
        stones = [-stone for stone in stones]
        heapq.heapify(stones)

        while len(stones) > 1:
            y = heapq.heappop(stones)
            x = heapq.heappop(stones)
        
            if x != y:
                heapq.heappush(stones, y - x)
        
        return -stones[0] if stones else 0


if __name__ == "__main__":
    o = Solution()

    # INPUT
    stones = [2,7,4,1,8,1]

    # OUTPUT
    result = o.lastStoneWeight(stones)
    print(result)
