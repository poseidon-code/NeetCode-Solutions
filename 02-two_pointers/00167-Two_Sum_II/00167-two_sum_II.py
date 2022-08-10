# 167: Two Sum II
# https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/


class Solution:
    # SOLUTION
    def twoSum(self, numbers: list[int], target: int) -> list[int]:
        l, r = 0, len(numbers) - 1

        while l<r:
            s = numbers[l] + numbers[r]
            if s==target: return [l+1, r+1]
            elif s<target: l+=1
            else: r-=1
        
        return []


if __name__ == "__main__":
    o = Solution()

    # INPUT
    numbers = [2,7,11,15]
    target = 9

    # OUTPUT
    result = o.twoSum(numbers, target)
    print(result)