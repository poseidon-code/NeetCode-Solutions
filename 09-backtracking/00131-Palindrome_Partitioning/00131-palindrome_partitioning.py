# 131: Palindrome Partitioning
# https://leetcode.com/problems/palindrome-partitioning



from typing import List

class Solution:
    # SOLUTION
    def partition(self, s: str) -> List[List[str]]:
        result = []
        current = []

        def isPalindrome(s: str, left: int, right: int) -> bool:
            while (left < right):
                if (s[left] != s[right]): return False
                left+=1
                right-=1

            return True


        def backtrack(start: int) -> None:
            if (start == len(s)):
                result.append(current.copy())
                return

            for i in range(start , len(s)):
                if isPalindrome(s, start, i):
                    str = s[start:i+1]
                    current.append(str)
                    backtrack(i+1)
                    current.pop()

        backtrack(0)
        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    s = "aab"

    # OUTPUT
    result = o.partition(s)

    print(result)
