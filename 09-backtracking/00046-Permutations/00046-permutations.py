# 46: Permutations
# https://leetcode.com/problems/permutations/


from typing import List

class Solution:
    # SOLUTION
    def permute(self, nums: List[int]) -> List[List[int]]:
        result = []

        def backtrack(start: int) -> None:
            if start == len(nums):
                result.append(nums[:])
                return

            for i in range(start , len(nums)):
                nums[i], nums[start] = nums[start], nums[i]
                backtrack(start+1)
                nums[i], nums[start] = nums[start], nums[i]

        backtrack(0)
        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums = [0,1]

    # OUTPUT
    result = o.permute(nums)

    print(result)
