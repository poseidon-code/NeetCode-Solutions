# 242: Valid Anagram
# https://leetcode.com/problems/valid-anagram/

from itertools import count


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t) : return False

        n = len(s)
        counts = [0]*26

        for i in range(n):
            counts[ord(s[i]) - ord('a')]+=1
            counts[ord(t[i]) - ord('a')]-=1

        for i in range(26):
            if counts[i] != 0 : return False
        
        return True


if __name__ == "__main__":
    o = Solution()

    # INPUT
    s: str = "rat"
    t: str = "cat"

    # OUTPUT
    result = o.isAnagram(s, t)
    print(result)
