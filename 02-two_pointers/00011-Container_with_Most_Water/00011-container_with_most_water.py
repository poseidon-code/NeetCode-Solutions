# 11: Container with Most Water
# https://leetcode.com/problems/container-with-most-water/


class Solution:
    # SOLUTION
    def maxArea(self, height: list[int]) -> int:
        l = 0
        r = len(height) - 1
        result = 0

        while l<r:
            result = max(result, (r-l)*min(height[l], height[r]))
            if height[l]<height[r]: l+=1
            else: r-=1
        
        return result


if __name__ == "__main__":
    o = Solution()

    # INPUT
    height: list[int] = [1,8,6,2,5,4,8,3,7]

    # OUTPUT
    result = o.maxArea(height)
    print(result)