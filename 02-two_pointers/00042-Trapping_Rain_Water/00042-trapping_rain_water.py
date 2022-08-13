# 42: Trapping Rain Water
# https://leetcode.com/problems/trapping-rain-water/


class Solution:
    # SOLUTION
    def trap(self, height: list[int]) -> int:
        l, r = 0, len(height) - 1
        maxLeft, maxRight = height[l], height[r]
        result = 0

        while l<r:
            if maxLeft<=maxRight:
                l+=1
                maxLeft = max(maxLeft, height[l])
                result += maxLeft - height[l]
            else:
                r-=1
                maxRight = max(maxRight, height[r])
                result += maxRight - height[r]

        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    height = [4,2,0,3,2,5]
    
    # OUTPUT
    result = o.trap(height)
    print(result)