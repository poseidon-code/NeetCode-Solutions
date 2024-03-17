# 703: Kth Largest Element In A Stream
# https://leetcode.com/problems/kth-largest-element-in-a-stream


import heapq


class KthLargest:
    def __init__(self, k: int, nums: list[int]) -> None:
        self.pq = nums
        self.size = k
        
        heapq.heapify(self.pq)
        while len(self.pq) > k:
            heapq.heappop(self.pq)

    def add(self, val: int) -> int:
        if len(self.pq) < self.size:
            heapq.heappush(self.pq, val)
        elif val > self.pq[0]:
            heapq.heapreplace(self.pq, val)
        return self.pq[0]



if __name__ == "__main__":

    # INPUT
    k = 3
    nums = [4,5,8,2]
    o = KthLargest(k, nums)

    # OPERATIONS
    print(o.add(3))
    print(o.add(5))
    print(o.add(10))
    print(o.add(9))
    print(o.add(4))
