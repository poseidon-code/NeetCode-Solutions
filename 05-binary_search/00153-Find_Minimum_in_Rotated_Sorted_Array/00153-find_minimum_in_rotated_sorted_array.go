// 153: Find Minimum in Rotated Sorted Array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

package main

import "fmt"

// SOLUTION
func findMin(nums []int) int {
    l := 0
    r := len(nums) - 1
    
    for (l<r && nums[l]>nums[r]) {
        m := l + (r-l) / 2

        if (nums[m] > nums[r]) {
            l = m + 1
        } else {
            r = m
        }
    }

    return nums[l]
}


func main() {
    // INPUT
    nums := []int{3,4,5,1,2}

    // OUTPUT
    result := findMin(nums)
    fmt.Println(result)
}