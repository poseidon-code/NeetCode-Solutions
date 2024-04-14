# 215: Kth Largest Element In An Array
# https://leetcode.com/problems/kth-largest-element-in-an-array


import heapq

class Solution:
    # SOLUTION
    def findKthLargest(self, nums: list[int], k: int) -> int:
        pq=[]
        count=0

        for num in nums:
            heapq.heappush(pq,num)
            count+=1
            if count > k:
                heapq.heappop(pq)
                count-=1

        return pq[0]


if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums = [3,2,1,5,6,4]
    k = 2

    # OUTPUT
    result = o.findKthLargest(nums, k)
    print(result)

