# 17: Letter Combinations Of A Phone Number
# https://leetcode.com/problems/letter-combinations-of-a-phone-number



from typing import List

class Solution:
    # SOLUTION
    def letterCombinations(self, digits: str) -> List[str]:
        if len(digits) == 0: 
            return []
        
        m = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "qprs",
            "8": "tuv",
            "9": "wxyz",
        }
        result = []
        current = ""

        def backtrack(index: int, current: str) -> None:
            if len(current) == len(digits):
                result.append(current)
                return
            
            for c in m[digits[index]]:
                backtrack(index+1, current + c)


        backtrack(0, current)
        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    digits = "23"

    # OUTPUT
    result = o.letterCombinations(digits)

    print(result)
