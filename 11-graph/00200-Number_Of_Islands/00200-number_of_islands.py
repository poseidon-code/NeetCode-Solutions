# 200: Number Of Islands
# https://leetcode.com/problems/number-of-islands/

from typing import List


class Solution:
    # SOLUTION
    def numIslands(self, grid: List[List[str]]) -> int:
        if not grid or not grid[0]:
            return 0

        result = 0
        visit = set()
        rows, cols = len(grid), len(grid[0])

        def dfs(i, j):
            if (
                i not in range(rows)
                or j not in range(cols)
                or grid[i][j] == "0"
                or (i, j) in visit
            ):
                return

            visit.add((i, j))
            dfs(i + 1, j);
            dfs(i, j + 1);
            dfs(i - 1, j);
            dfs(i, j - 1);

        for i in range(rows):
            for j in range(cols):
                if grid[i][j] == "1" and (i, j) not in visit:
                    result += 1
                    dfs(i, j)

        return result


if __name__ == "__main__":
    o = Solution()

    # INPUT
    grid =[
        ['1','1','1','1','0'],
        ['1','1','0','1','0'],
        ['1','1','0','0','0'],
        ['0','0','0','0','0']
    ]

    # OUTPUT
    result = o.numIslands(grid)
    print(result)
