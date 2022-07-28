# 347: Top K Frequent Elements
# https://leetcode.com/problems/top-k-frequent-elements/

class Solution:
    # SOLUTION
    def topKFrequent(self, nums: list[int], k: int) -> list[int]:
        count = {}
        buckets = [[] for _ in range(len(nums) + 1)]

        for n in nums:
            count[n] = 1 + count.get(n, 0)
        for n, c in count.items():
            buckets[c].append(n)

        result = []
        for i in range(len(buckets)-1, 0, -1):
            for n in buckets[i]:
                result.append(n)
            if len(result) == k: break

        return result


if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums = [1,1,1,2,2,3]
    k = 2

    # OUTPUT
    result = o.topKFrequent(nums, k)
    print(result)
