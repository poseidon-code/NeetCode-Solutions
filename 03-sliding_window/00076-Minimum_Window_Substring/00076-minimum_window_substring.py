# 76: Minimum Window Substring
# https://leetcode.com/problems/minimum-window-substring/


class Solution:
    # SOLUTION
    def minWindow(self, s: str, t: str) -> str:
        m = {}
        for i in range(len(t)): m[t[i]] = 1 + m.get(t[i], 0)

        i = 0
        j = 0
        counter = len(t)
        
        minStart = 0
        minLength = len(s)+1
        
        while j < len(s):
            if m.get(s[j]) != None: 
                m[s[j]]-=1 
                counter-=1
            j+=1

            while counter==0:
                if j-i < minLength:
                    minStart = i
                    minLength = j-i
                
                if m.get(s[i]) != None:
                    m[s[i]]+=1
                    counter+=1
                i+=1
        
        if (minLength != len(s)):
            return s[minStart:minStart+minLength]
        return ""



if __name__ == "__main__":
    o = Solution()

    # INPUT
    s: str = "ADOBECODEBANC"
    t: str = "ABC"

    # OUTPUT
    result = o.minWindow(s, t)
    print(result)