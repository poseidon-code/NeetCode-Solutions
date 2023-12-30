# 212: Word Search Ii
# https://leetcode.com/problems/word-search-ii/



from typing import DefaultDict, List

class TrieNode:
    def __init__(self):
        self.children = DefaultDict(TrieNode)
        self.word = None

    def addWord(self, word):
        cur = self
        for c in word:
            cur = cur.children[c]
        cur.word = word


class Solution:
    # SOLUTION
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        m, n = len(board), len(board[0])
        DIR = [0, 1, 0, -1, 0]
        trieNode = TrieNode()
        result = []
        for word in words:
            trieNode.addWord(word)

        def dfs(r, c, cur):
            if r<0 or r==m or c<0 or c==n or board[r][c] not in cur.children: return
            
            orgChar = board[r][c]
            cur = cur.children[orgChar]
            board[r][c] = '#'
            
            if cur.word != None:
                result.append(cur.word)
                cur.word = None
            
            for i in range(4): dfs(r + DIR[i], c + DIR[i + 1], cur)
            board[r][c] = orgChar

        for r in range(m):
            for c in range(n):
                dfs(r, c, trieNode)
        
        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    board = [
        ["o","a","a","n"],
        ["e","t","a","e"],
        ["i","h","k","r"],
        ["i","f","l","v"]
    ]
    words = ["oath","pea","eat","rain"]

    # OUTPUT
    print(o.findWords(board, words))
