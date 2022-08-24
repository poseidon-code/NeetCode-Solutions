# 22: Generate Parentheses
# https://leetcode.com/problems/generate-parentheses/


class Solution:
    # SOLUTION    
    def generateParentheses(self, n: int) -> list[str]:
        result: list[str] = []
        self.generate(n, 0, 0, "", result)
        return result

    def generate(self, n: int, open: int, close: int, s: str, result) -> None:
        if open==n and close==n:
            result.append(s)
            return
        if open < n:
            self.generate(n, open+1, close, s+'(', result)
        if open > close:
            self.generate(n, open, close+1, s+')', result)



if __name__ == "__main__":
    o = Solution()

    # INPUT
    n: int = 3

    # OUTPUT
    result = o.generateParentheses(n)
    print(result)