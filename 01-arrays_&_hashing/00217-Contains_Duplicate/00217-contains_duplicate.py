# 217: Contains Duplicate
# https://leetcode.com/problems/contains-duplicate/

class Solution:
    # SOLUTION
    def containsDuplicate(self, nums: list[int]) -> bool:
        return len(nums) > len(set(nums))

if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums: list[int] = [1,2,3,1]

    # OUTPUT
    result = o.containsDuplicate(nums)
    print(result)
