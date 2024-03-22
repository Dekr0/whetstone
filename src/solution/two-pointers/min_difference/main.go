package main

import (
	"slices"
)

// two pointer + sliding window
func solutionA(nums []int, k int) int {
    if k == 1 || len(nums) == 1 {
        return 0
    }

    slices.Sort(nums)

    r := -1
    for i := 0; i + k <= len(nums); i++ {
        if r == -1 {
            r = nums[i + k - 1] - nums[i]
        } else {
            r = min(r, nums[i + k - 1] - nums[i])
        }
    }

    return r
}
