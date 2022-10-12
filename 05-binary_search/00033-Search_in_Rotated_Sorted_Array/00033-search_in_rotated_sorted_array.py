# 33. Search in Rotated Sorted Array
# https://leetcode.com/problems/search-in-rotated-sorted-array/


class Solution:
    # SOLUTION
    def search(self, nums: list[int], target: int) -> int:
        l, h = 0, len(nums)
        
        while l < h:
            m = (l+h) // 2
            if target < nums[0] < nums[m]:
                l = m+1
            elif target >= nums[0] > nums[m]:
                h = m
            elif nums[m] < target:
                l = m+1
            elif nums[m] > target:
                h = m
            else:
                return m

        return -1



if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums: list[int] = [4,5,6,7,0,1,2]
    target: int = 0

    # OUTPUT
    result = o.search(nums, target)
    print(result)
