// 4: Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func findMedianSortedArrays(nums1, nums2 []int) float64 {
    N1 := len(nums1)
    N2 := len(nums2)
    if N1 < N2 { return findMedianSortedArrays(nums2, nums1) }
    
    l, h := 0, N2 * 2
    
    for (l <= h) {
        mid2 := (l + h) / 2
        mid1 := N1 + N2 - mid2

        L1 := func()int{if mid1==0 {return math.MinInt16} else {return nums1[(mid1-1)/2]}}()
        L2 := func()int{if mid2==0 {return math.MinInt16} else {return nums2[(mid2-1)/2]}}()
        R1 := func()int{if mid1==N1*2 {return math.MaxInt16} else {return nums1[(mid1)/2]}}()
        R2 := func()int{if mid2==N2*2 {return math.MaxInt16} else {return nums2[(mid2)/2]}}()

        if L1 > R2 {
            l = mid2 + 1
        } else if L2 > R1 {
            h = mid2 - 1
        } else {
            return (math.Max(float64(L1), float64(L2)) + math.Min(float64(R1), float64(R2))) / 2
        }
    }

    return -1;
}


func main() {
    // INPUT
    nums1 := []int{1,3}
    nums2 := []int{2}

    // OUTPUT
    result := findMedianSortedArrays(nums1, nums2)
    fmt.Println(result)
}
