# 51: N Queens
# https://leetcode.com/problems/n-queens


from typing import List

class Solution:
    # SOLUTION
    def solveNQueens(self, n: int) -> List[List[str]]:
        result = []
        nQueens = [("." * n)] * n

        def isValid(row: int, col: int) -> bool:
            for i in range(row):
                if nQueens[i][col] == 'Q':
                    return False

            i, j = row-1, col-1
            while i>=0 and j>=0:
                if nQueens[i][j] == 'Q':
                    return False
                i -= 1
                j -= 1

            i, j = row-1, col+1
            while i>=0 and j<n:
                if nQueens[i][j] == 'Q':
                    return False
                i -= 1
                j += 1

            return True


        def backtrack(row: int) -> None:
            if (row == n):
                result.append(nQueens.copy())
                return

            for col in range(n):
                if isValid(row, col):
                    r =  nQueens[row]
                    r = r[0:col] + "Q" + r[col+1:]
                    nQueens[row] = r
                    backtrack(row+1)
                    r = r[0:col] + "." + r[col+1:]
                    nQueens[row] = r

        backtrack(0)
        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    n = 4

    # OUTPUT
    result = o.solveNQueens(n)

    print(result)
