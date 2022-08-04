# 36: Valid Sudoku
# https://leetcode.com/problems/valid-sudoku/


class Solution:
    # SOLUTION
    def isValidSudoku(self, board: list[list[str]]) -> bool:
        used1, used2, used3 = [[0 for _ in range(9)] for _ in range(9)], [[0 for _ in range(9)] for _ in range(9)], [[0 for _ in range(9)] for _ in range(9)]

        for i in range(len(board)):
            for j in range(len(board[i])):
                if board[i][j] != ".":
                    num = int(board[i][j]) - 1
                    k = i//3 * 3 + j//3

                    if used1[i][num] or used2[j][num] or used3[k][num]:
                        return False

                    used1[i][num] = used2[j][num] = used3[k][num] = 1
        
        return True


if __name__ == "__main__":
    o = Solution()

    # INPUT
    board: list[list[str]] = [
        ["5","3",".",".","7",".",".",".","."],
        ["6",".",".","1","9","5",".",".","."],
        [".","9","8",".",".",".",".","6","."],
        ["8",".",".",".","6",".",".",".","3"],
        ["4",".",".","8",".","3",".",".","1"],
        ["7",".",".",".","2",".",".",".","6"],
        [".","6",".",".",".",".","2","8","."],
        [".",".",".","4","1","9",".",".","5"],
        [".",".",".",".","8",".",".","7","9"]
    ]

    # OUTPUT
    result = o.isValidSudoku(board)
    print(result)