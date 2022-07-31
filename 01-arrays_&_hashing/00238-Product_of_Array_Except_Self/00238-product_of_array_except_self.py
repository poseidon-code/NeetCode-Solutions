# 238 : Product of Array Except Self
# https://leetcode.com/problems/product-of-array-except-self/


class Solution:
    # SOLUTION
    def productExceptSelf(self, nums: list[int]) -> list[int]:
        result, suf, pre = [1]*len(nums), 1, 1

        for i in range(len(nums)):
            result[i] *= pre
            pre *= nums[i]
            result[-1-i] *= suf
            suf *= nums[-1-i]
        
        return result


if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums: list[int] = [1,2,3,4]

    # OUTPUT
    result = o.productExceptSelf(nums)
    print(result)