// 287: Find the Duplicate Number
// https://leetcode.com/problems/find-the-duplicate-number/

package main

import "fmt"

// SOLUTION
func findDuplicate(nums []int) int {
    slow, fast := nums[0], nums[nums[0]]
        
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }
    
    slow = 0
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }

    return slow
}

func main() {
    // INPUT
    nums := []int{1,3,4,2,2}

    // OUTPUT
    result := findDuplicate(nums)
    fmt.Println(result)
}