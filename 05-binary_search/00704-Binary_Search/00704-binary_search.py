# 704: Binary Search
# https://leetcode.com/problems/binary-search/


class Solution:
    # SOLUTION
    def search(self, nums: list[int], target: int) -> int:
        l, h = 0, len(nums) - 1

        while l <= h:
            m = l + (h-l)//2
            
            if nums[m] < target:
                l = m+1
            elif nums[m] > target:
                h = m-1
            else:
                return m

        return -1



if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums: list[int] = [-1,0,3,5,9,12]
    target: int = 9

    # OUTPUT
    result = o.search(nums, target)
    print(result)