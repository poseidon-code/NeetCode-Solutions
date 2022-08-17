# 567: Permutation in String
# https://leetcode.com/problems/permutation-in-string/


class Solution:
    # SOLUTION
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2): return False
        
        s1Count = [0]*26
        for c in s1: s1Count[ord(c)-ord('a')]+=1
        s2Count = [0]*26
        i, j = 0, 0
        
        while j < len(s2):
            s2Count[ord(s2[j])-ord('a')]+=1
            
            if j-i+1 == len(s1):
                if s1Count == s2Count:
                    return True
            if j-i+1 < len(s1):
                j+=1
            else:
                s2Count[ord(s2[i])-ord('a')]-=1
                i+=1
                j+=1


        return False



if __name__ == "__main__":
    o = Solution()

    # INPUT
    s1: str = "ab"
    s2: str = "eidbaooo"

    # OUTPUT
    result = o.checkInclusion(s1, s2)
    print(result)