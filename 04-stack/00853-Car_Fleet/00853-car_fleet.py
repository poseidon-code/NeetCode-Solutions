# 853: Car Fleet
# https://leetcode.com/problems/car-fleet/


import operator


class Solution:
    class car:
        def __init__(self, p: int, s: int):
            self.p = p
            self.s = s

    # SOLUTION
    def carFleet(self, target: int, position: list[int], speed: list[int]) -> int:
        cars: list[car] = []
        n = len(position)
        for i in range(n):
            cars.append(self.car(position[i], speed[i]))

        cars.sort(key=operator.attrgetter('p'))

        s: list[float] = []
        for i in range(n):
            time = (target - cars[i].p) / cars[i].s
            while len(s)!=0 and time >= s[-1]: s.pop()
            s.append(time)

        return len(s)



if __name__ == "__main__":
    o = Solution()

    # INPUT
    target: int = 12
    position: list[int] = [10,8,0,5,3]
    speed: list[int] = [2,4,1,1,3]

    # OUTPUT
    result = o.carFleet(target, position, speed)
    print(result)