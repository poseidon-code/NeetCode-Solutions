# 695: Max Area Of Island
# https://leetcode.com/problems/max-area-of-island

from typing import List


class Solution:
    # SOLUTION
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        visit = set()

        def dfs(i: int, j: int):
            if (i < 0 or i == m or j < 0 or j == n or grid[i][j] == 0 or (i, j) in visit):
                return 0

            visit.add((i, j))

            return \
                1 + dfs(i + 1, j) \
                + dfs(i - 1, j) \
                + dfs(i, j + 1) \
                + dfs(i, j - 1)


        result = 0
        for r in range(m):
            for c in range(n):
                result = max(result, dfs(r, c))

        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    grid = [
        [0,0,1,0,0,0,0,1,0,0,0,0,0],
        [0,0,0,0,0,0,0,1,1,1,0,0,0],
        [0,1,1,0,1,0,0,0,0,0,0,0,0],
        [0,1,0,0,1,1,0,0,1,0,1,0,0],
        [0,1,0,0,1,1,0,0,1,1,1,0,0],
        [0,0,0,0,0,0,0,0,0,0,1,0,0],
        [0,0,0,0,0,0,0,1,1,1,0,0,0],
        [0,0,0,0,0,0,0,1,1,0,0,0,0]
    ]

    # OUTPUT
    result = o.maxAreaOfIsland(grid)
    print(result)
