// 4: Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/

class Solution {
    // SOLUTION
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int N1 = nums1.length;
        int N2 = nums2.length;
        if (N1 < N2) return findMedianSortedArrays(nums2, nums1);
        
        int l = 0, h = N2 * 2;
        
        while (l <= h) {
            int mid2 = (l + h) / 2;
            int mid1 = N1 + N2 - mid2;
            

            double L1 = (mid1 == 0) ? Integer.MIN_VALUE : nums1[(mid1-1)/2];
            double L2 = (mid2 == 0) ? Integer.MIN_VALUE : nums2[(mid2-1)/2];
            double R1 = (mid1 == N1 * 2) ? Integer.MAX_VALUE : nums1[(mid1)/2];
            double R2 = (mid2 == N2 * 2) ? Integer.MAX_VALUE : nums2[(mid2)/2];

            if (L1 > R2) l = mid2 + 1;
            else if (L2 > R1) h = mid2 - 1;
            else return (Math.max(L1,L2) + Math.min(R1, R2)) / 2;
        }

        return -1;
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        
        // INPUT
        int[] nums1 = {1,3};
        int[] nums2 = {2};

        // OUTPUT
        var result = o.findMedianSortedArrays(nums1, nums2);
        System.out.println(result);
    }
}