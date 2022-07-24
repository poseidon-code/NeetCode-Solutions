# 1: Two Sum
# https://leetcode.com/problems/two-sum/

class Solution:
    # SOLUTION
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        map = {}
        for i, n in enumerate(nums):
            m = target - n
            if m in map:
                return [map[m], i]
            else:
                map[n] = i


if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums: list[int] = [3,2,4]
    target: int = 6

    # OUTPUT
    result = o.twoSum(nums, target)
    print(result)
