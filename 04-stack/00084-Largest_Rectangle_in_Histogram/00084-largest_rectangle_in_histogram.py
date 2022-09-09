# 84: Largest Rectangle in Histogram
# https://leetcode.com/problems/largest-rectangle-in-histogram/


class Solution:
    def largestRectangleAre(self, heights: list[int]) -> int:
        stack = []
        l = len(heights)
        result = 0

        i=0
        while i<=l:
            h = 0 if i==l else heights[i]

            if len(stack)==0 or h>=heights[stack[-1]]:
                stack.append(i)
            else:
                tp = stack[-1]
                stack.pop()
                result = max(result, heights[tp]*(i if len(stack)==0 else i-1-stack[-1]))
                i-=1
            i+=1

        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    heights: list[int] = [2,1,5,6,2,3]

    # OUTPUT
    result = o.largestRectangleAre(heights)
    print(result)