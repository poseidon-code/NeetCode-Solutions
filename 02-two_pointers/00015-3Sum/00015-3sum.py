# 15: 3Sum
# https://leetcode.com/problems/3sum/


class Solution:
    # SOLUTION
    def threeSum(self, nums: list[int]) -> list[list[int]]:
        nums.sort()
        result: list[list[int]] = []

        for i in range(len(nums)):
            if i>0 and nums[i]==nums[i-1]: continue

            l = i+1
            r = len(nums) - 1

            while l<r:
                s = nums[i] + nums[l] + nums[r]
                if s>0: r-=1
                elif s<0: l+=1
                else:
                    result.append([nums[i], nums[l], nums[r]])
                    while nums[l]==nums[l+1]: l+=1
                    while nums[r]==nums[r-1]: r-=1
                    l+=1
                    r-=1
            
        return result



if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums: list[int] = [-1,0,1,2,-1,-4]

    # OUTPUT
    result = o.threeSum(nums)
    print(result)