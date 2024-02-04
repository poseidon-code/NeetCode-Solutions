# 90: Subsets Ii
# https://leetcode.com/problems/subsets-ii


from typing import List

class Solution:
    # SOLUTION
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        result = []
        current = []

        def backtrack(start: int) -> None:
            result.append(current.copy())

            for i in range(start , len(nums)):
                if i>start and nums[i] == nums[i-1] : continue
                current.append(nums[i])
                backtrack(i+1)
                current.pop()

        backtrack(0)
        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums = [1,2,2]

    # OUTPUT
    result = o.subsetsWithDup(nums)

    print(result)
