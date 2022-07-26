# 49: Group Anagrams
# https://leetcode.com/problems/group-anagrams/

from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: list[str]) -> list[list[str]]:
        tmap = defaultdict(list)
        
        for s in strs:
            tmap[str(sorted(s))].append(s)
        
        return list(tmap.values())


if __name__ == "__main__":
    o = Solution()

    # INPUT
    strs = ["eat","tea","tan","ate","nat","bat"]

    # OUTPUT
    result = o.groupAnagrams(strs)
    print(result)
