# 973: K Closest Points To Origin
# https://leetcode.com/problems/k-closest-points-to-origin


import heapq

class Solution:
    # SOLUTOIN
    def kClosest(self, points: list[list[int]], k: int) -> list[list[int]]:
        pq = []
        for x, y in points:
            dist = (x ** 2) + (y ** 2)
            pq.append((dist, x, y))

        heapq.heapify(pq)
        result = []

        for _ in range(k):
            _, x, y = heapq.heappop(pq)
            result.append((x, y))

        return result


if __name__ == "__main__":
    o = Solution()

    # INPUT
    points = [[1,3],[-2,2]]
    k = 1

    # OUTPUT
    result = o.kClosest(points, k)
    print(result)
