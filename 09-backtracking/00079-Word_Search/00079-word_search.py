# 79: Word Search
# https://leetcode.com/problems/word-search


from typing import List

class Solution:
    # SOLUTION
    def exist(self, board: List[List[str]], word: str) -> bool:
        def backtrack(board: List[List[str]], i: int, j: int, word: str):
            if len(word) == 0 : return True
            if i<0 or i>=len(board) or j<0 or j>= len(board[0]) or board[i][j]!=word[0] : return False

            c = board[i][j]
            board[i][j] = '*'
            s = word[1:]

            result =   backtrack(board, i-1, j, s) \
                    or backtrack(board, i+1, j, s) \
                    or backtrack(board, i, j-1, s) \
                    or backtrack(board, i, j+1, s)
            board[i][j] = c

            return result


        for i in range(len(board)):
            for j in range(len(board[0])):
                if backtrack(board, i, j, word):
                    return True
        
        return False



if __name__ == "__main__":
    o = Solution()

    # INPUT
    board = [
        ["A","B","C","E"],
        ["S","F","C","S"],
        ["A","D","E","E"]
    ]
    word = "ABCCED"

    # OUTPUT
    result = o.exist(board, word)

    print(result)
