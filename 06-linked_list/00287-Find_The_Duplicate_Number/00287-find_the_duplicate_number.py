# 287: Find the Duplicate Number
# https://leetcode.com/problems/find-the-duplicate-number/


class Solution:
    # SOLUTION
    def findDuplicate(self, nums: list[int]) -> int:
        slow, fast = nums[0], nums[nums[0]]
        
        while slow != fast:
            slow = nums[slow]
            fast = nums[nums[fast]]
        
        slow = 0
        while slow != fast:
            slow = nums[slow]
            fast = nums[fast]

        return slow


if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums: list[int] = [1,3,4,2,2]
    
    # OUTPUT
    result = o.findDuplicate(nums)
    print(result)