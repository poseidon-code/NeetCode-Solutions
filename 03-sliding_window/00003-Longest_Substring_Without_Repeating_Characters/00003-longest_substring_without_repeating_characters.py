# 3: Longest Substring Without Repeating Characters
# https://leetcode.com/problems/longest-substring-without-repeating-characters/


class Solution:
    # SOLUTION
    def lengthOfLongestSubstring(self, s: str) -> int:
        letters = set()
        i, j = 0, 0
        result = 0

        for _ in range(len(s)):
            while s[j] in letters:
                letters.remove(s[i])
                i += 1

            letters.add(s[j])
            result = max(result, j-i+1)
            j += 1

        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    s: str = "abcabcbb"

    # OUTPUT
    result = o.lengthOfLongestSubstring(s)
    print(result)