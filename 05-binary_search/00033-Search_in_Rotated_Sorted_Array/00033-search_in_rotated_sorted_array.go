// 33. Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/

package main

import "fmt"

// SOLUTION
func search(nums []int, target int) int {
    l, h := 0, len(nums)

    for (l < h) {
        m := (l+h) / 2
        if (target < nums[0] && nums[0] < nums[m]) {
            l = m+1
        } else if (target >= nums[0] && nums[0] > nums[m]) {
            h = m
        } else if (nums[m] < target) {
            l = m+1
        } else if (nums[m] > target) {
            h = m
        } else {
            return m
        }
    }

    return -1
}


func main() {
    // INPUT
    nums := []int{4,5,6,7,0,1,2}
    target := 0
    
    // OUTPUT
    result := search(nums, target)
    fmt.Println(result)
}