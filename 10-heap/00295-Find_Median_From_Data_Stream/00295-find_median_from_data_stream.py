# 295: Find Median From Data Stream
# https://leetcode.com/problems/find-median-from-data-stream

import heapq

# SOLUTION
class MedianFinder:
    def __init__(self):
        self.lower, self.higher = [], []

    def addNum(self, num: int) -> None:
        if self.higher and num > self.higher[0]:
            heapq.heappush(self.higher, num)
        else:
            heapq.heappush(self.lower, -1 * num)

        if len(self.lower) > len(self.higher) + 1:
            value = -1 * heapq.heappop(self.lower)
            heapq.heappush(self.higher, value)

        if len(self.higher) > len(self.lower) + 1:
            value = heapq.heappop(self.higher)
            heapq.heappush(self.lower, -1 * value)

    def findMedian(self) -> float:
        if len(self.lower) > len(self.higher):
            return -1 * self.lower[0]
        elif len(self.higher) > len(self.lower):
            return self.higher[0]
        return (-1 * self.lower[0] + self.higher[0]) / 2.0


if __name__ == "__main__":
    o = MedianFinder()

    # OPERATIONS
    o.addNum(1)
    o.addNum(2)
    print(o.findMedian())
    o.addNum(3)
    print(o.findMedian())
