# 239: Sliding Window Maximum
# https://leetcode.com/problems/sliding-window-maximum/

import collections

class Solution:
    # SOLUTION
    def maxSlidingWindow(self, nums: list[int], k: int) -> list[int]:
        dq = collections.deque()
        result = []
        
        for i, n in enumerate(nums):
            while dq and nums[dq[-1]] < n:
                dq.pop()
            
            dq += i,
            
            if dq[0] == i - k:
                dq.popleft()
            
            if i >= k - 1:
                result += nums[dq[0]],
        
        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums: list[int] = [1,3,-1,-3,5,3,6,7]
    k: int = 3

    # OUTPUT
    result = o.maxSlidingWindow(nums, k);
    print(result)