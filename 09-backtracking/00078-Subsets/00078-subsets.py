# 78: Subsets
# https://leetcode.com/problems/subsets/

from typing import List

class Solution:
    # SOLUTION
    def subsets(self, nums: List[int]) -> List[List[int]]:
        result = []
        current = []

        def backtrack(start: int) -> None:
            result.append(current.copy())

            for i in range(start , len(nums)):
                current.append(nums[i])
                backtrack(i+1)
                current.pop()

        backtrack(0)
        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums = [1,2,3]

    # OUTPUT
    result = o.subsets(nums)

    print(result)
