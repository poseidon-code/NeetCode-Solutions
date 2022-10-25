# 153: Find Minimum in Rotated Sorted Array
# https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/


class Solution:
    # SOLUTION
    def findMin(self, nums: list[int]) -> int:
        l = 0
        r = len(nums) - 1
        
        while (l<r and nums[l]>nums[r]):
            m = l + (r-l) // 2

            if (nums[m] > nums[r]):
                l = m + 1
            else:
                r = m

        return nums[l]



if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums: list[int] = [3,4,5,1,2]

    # OUTPUT
    result = o.findMin(nums)
    print(result)