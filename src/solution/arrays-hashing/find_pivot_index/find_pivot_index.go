package findpivotindex

func FindPivotIndex(nums []int) int {
    left := make([]int, len(nums), len(nums))
    right := make([]int, len(nums), len(nums))
    l_sum := 0
    r_sum := 0
    for i := 0; i < len(nums); i++ {
        if i == 0 {
            continue
        } else {
            l_sum += nums[i - 1]
            r_sum += nums[len(nums) - 1 - i + 1]
            left[i] = l_sum
            right[len(nums) - 1 - i] = r_sum
        }
    }
    for i := 0; i < len(nums); i++ {
        if left[i] == right[i] {
            return i
        }
    }

    return -1
}

func FindPivotIndexOpt(nums []int) int {
    // Optimization => O(1) space => state
    // goal => break at the point where left sum == right sum
    // how to obtain left sum at each point? accum. and traverse from 
    // left to right
    // how to obtain right sum at each point? Total - left sum at each point
    sum := 0
    for _, n := range nums { sum += n }
    lsum := 0
    for i := 0; i < len(nums); i++ {
        if lsum == sum - lsum - nums[i] {
            return i
        }
        lsum += nums[i]
    }

    return -1
}

