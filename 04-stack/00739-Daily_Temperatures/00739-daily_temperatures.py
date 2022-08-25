# 739: Daily Temperstures
# https://leetcode.com/problems/daily-temperatures/


class Solution:
    # SOLUTION
    def dailyTemperatures(self, temperatures: list[int]) -> list[int]:
        s = []
        result = [0]*len(temperatures)

        for i in range(len(temperatures)):
            while len(s)!=0 and temperatures[i]>temperatures[s[-1]]:
                result[s[-1]] = i - s[-1]
                s.pop()
            s.append(i)

        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    temperatures: list[int] = [73,74,75,71,69,72,76,73]

    # OUTPUT
    result = o.dailyTemperatures(temperatures)
    print(result)