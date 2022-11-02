// 4: Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/


#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;


class Solution {
public:
    // SOLUITON
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int N1 = nums1.size();
        int N2 = nums2.size();
        if (N1 < N2) return findMedianSortedArrays(nums2, nums1);
        
        int l = 0, h = N2 * 2;
        
        while (l <= h) {
            int mid2 = (l + h) / 2;
            int mid1 = N1 + N2 - mid2;

            double L1 = (mid1 == 0) ? INT16_MIN : nums1[(mid1-1)/2];
            double L2 = (mid2 == 0) ? INT16_MIN : nums2[(mid2-1)/2];
            double R1 = (mid1 == N1 * 2) ? INT16_MAX : nums1[(mid1)/2];
            double R2 = (mid2 == N2 * 2) ? INT16_MAX : nums2[(mid2)/2];

            if (L1 > R2) l = mid2 + 1;
            else if (L2 > R1) h = mid2 - 1;
            else return (max(L1,L2) + min(R1, R2)) / 2;
        }

        return -1;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> nums1 = {1,3};
    vector<int> nums2 = {2};

    // OUTPUT
    auto result = o.findMedianSortedArrays(nums1, nums2);
    cout << result << endl;

    return 0;
}