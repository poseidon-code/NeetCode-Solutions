# 20: Valid Parentheses
# https://leetcode.com/problems/valid-parentheses/


class Solution:
    # SOLUTION
    def isValid (self, s: str) -> bool:
        parenthesis = []

        for c in s:
            if c=='{':
                parenthesis.append('}')
            elif c=='[':
                parenthesis.append(']')
            elif c=='(':
                parenthesis.append(')')
            else:
                if len(parenthesis)==0 or c!=parenthesis.pop():
                    return False
        
        return len(parenthesis)==0


if __name__ == "__main__":
    o = Solution()

    # INPUT
    s = "()[]{}"

    # OUTPUT
    result = o.isValid(s)
    print(result)