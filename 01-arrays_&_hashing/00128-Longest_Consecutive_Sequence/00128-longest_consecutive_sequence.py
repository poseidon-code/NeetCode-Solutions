# 00128: Longest Consecutive Sequence
# https://leetcode.com/problems/longest-consecutive-sequence/


class Solution:
    def longestConsecutive(self, nums: list[int]) -> int:
        numSet = set(nums)
        result = 0

        for i in nums:
            if (i - 1) not in numSet:
                length = 1
                while (i + length) in numSet:
                    length += 1
                result = max(length, result)
        
        return result


if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums: list[int] = [100,4,200,1,3,2]

    # OUTPUT
    result = o.longestConsecutive(nums)
    print(result)