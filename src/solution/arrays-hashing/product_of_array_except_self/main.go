package main

import (
	"fmt"
	"runtime/debug"
)


func productExceptSelf(nums []int) []int {
    n := len(nums)

    left := make([]int, n)
    right := make([]int, n)
    result := make([]int, n)
    left[0] = 1
    right[n - 1] = 1

    lp := nums[0]
    rp := nums[n - 1]
    for i := 1; i < n; i++ {
        left[i] = lp
        right[n - 1 - i] = rp
        lp *= nums[i]
        rp *= nums[n - 1 - i]
    }

    for i, v := range left {
        result[i] = v * right[i]
    }
    
    return result
}

func solution(nums []int) []int {
    debug.SetGCPercent(0)
    result := make([]int, len(nums))
    result[0] = 1
    p := nums[0]
    for i := 1; i < len(nums); i++ {
        result[i] = p 
        p *= nums[i]
    }
    p = nums[len(nums) - 1]
    for i := len(nums) - 2; i >= 0; i-- {
        result[i] *= p
        p *= nums[i]
    }

    return result
} 

func main() {
    fmt.Println(solution([]int{1, 2, 3, 4}))
}
