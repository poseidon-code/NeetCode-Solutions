# 40: Combination Sum Ii
# https://leetcode.com/problems/combination-sum-ii


from typing import List

class Solution:
    # SOLUTION
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        result = []
        current = []

        def backtrack(target: int, sum: int, start: int) -> None:
            if start > target : return

            if sum == target:
                result.append(current.copy())
                return

            for i in range(start, len(candidates)):
                if i>start and candidates[i] == candidates[i-1]: continue
                current.append(candidates[i])
                backtrack(target, sum + candidates[i], i+1)
                current.pop()

        backtrack(target, 0, 0)
        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    candidates = [2,5,2,1,2]
    target = 5

    # OUTPUT
    result = o.combinationSum2(candidates, target)

    print(result)
