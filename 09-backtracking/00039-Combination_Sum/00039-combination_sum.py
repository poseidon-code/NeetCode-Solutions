# 39: Combination Sum
# https://leetcode.com/problems/combination-sum/


from typing import List

class Solution:
    # SOLUTION
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        result = []
        current = []

        def backtrack(target: int, start: int) -> None:
            if start >= len(candidates) or target < 0 : return

            if target == 0:
                result.append(current.copy())
                return

            current.append(candidates[start])
            backtrack(target - candidates[start], start)
            current.pop()
            backtrack(target, start + 1)

        backtrack(target, 0)
        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    candidates = [2,3,6,7]
    target = 7

    # OUTPUT
    result = o.combinationSum(candidates, target)

    print(result)
