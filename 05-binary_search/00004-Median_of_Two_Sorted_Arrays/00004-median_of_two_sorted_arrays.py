# 4: Median of Two Sorted Arrays
# https://leetcode.com/problems/median-of-two-sorted-arrays/


INT16_MIN = -(2**16)
INT16_MAX = 2**16

class Solution:
    # SOLUTION
    def findMedianSortedArrays(self, nums1: list[int], nums2: list[int]) -> float:
        N1 = len(nums1)
        N2 = len(nums2)
        if N1 < N2: return self.findMedianSortedArrays(nums2, nums1)

        l, h = 0, N2 * 2

        while (l <= h):
            mid2 = (l + h) // 2
            mid1 = N1 + N2 - mid2

            L1 = INT16_MIN if mid1 == 0 else nums1[(mid1-1)//2]
            L2 = INT16_MIN if mid2 == 0 else nums2[(mid2-1)//2]
            R1 = INT16_MAX if mid1 == N1*2 else nums1[(mid1)//2]
            R2 = INT16_MAX if mid2 == N2*2 else nums2[(mid2)//2]

            if L1 > R2: l = mid2 + 1
            elif L2 > R1: h = mid2 - 1
            else: return (max(L1, L2) + min(R1, R2)) / 2
        
        return -1



if __name__ == "__main__":
    o = Solution()

    # INPUT
    nums1 = [1,3]
    nums2 = [2]

    # OUTPUT
    result = o.findMedianSortedArrays(nums1, nums2)
    print(result)
