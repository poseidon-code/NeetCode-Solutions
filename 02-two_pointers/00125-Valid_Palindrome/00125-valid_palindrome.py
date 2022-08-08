# 125: Valid Palindrome
# https://leetcode.com/problems/valid-palindrome/

class Solution:
    # SOLUTION
    def isPalindrome(self, s: str) -> bool:
        l = 0
        r = len(s) - 1

        while l<r:
            while not s[l].isalnum() and l<r: l+=1
            while not s[r].isalnum() and l<r: r-=1
            if s[l].lower() != s[r].lower(): return False

            l+=1
            r-=1

        return True



if __name__ == "__main__":
    o = Solution()

    # INPUT
    s: str = "race a car"

    # OUTPUT
    result = o.isPalindrome(s)
    print(result)