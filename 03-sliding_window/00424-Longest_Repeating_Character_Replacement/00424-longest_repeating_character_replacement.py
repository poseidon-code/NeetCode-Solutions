# 424: Longest Repeating Character Replacement
# https://leetcode.com/problems/longest-repeating-character-replacement/


class Solution:
    # SOLUTION
    def characterReplacement(self, s: str, k: int) -> int:
        count = [0]*26
        maxCount = 0
        i, j = 0, 0

        result = 0

        while j<len(s):
            count[ord(s[j])-ord('A')]+=1
            maxCount = max(maxCount, count[ord(s[j])-ord('A')])

            if j-i+1-maxCount > k:
                count[ord(s[i])-ord('A')]-=1
                i+=1
            result = max(result, j-i+1)
            j+=1
        
        return result




if __name__ == "__main__":
    o = Solution()

    # INPUT
    s: str = "ABAB"
    k: int = 2

    # OUTPUT
    result = o.characterReplacement(s, k)
    print(result)