package removeelement

func RemoveElement(nums []int, val int) int {
    // alternate version of pivoting
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[j] = nums[i] 
            // since elements are removing, we can savely override
            j++
        }
    }
    return j
}
