package rangesumqueryimmutable

type NumArray struct {
    nums []int
    prefixs []int
}

func Constructor(nums []int) NumArray {
    prefixs := make([]int, len(nums), len(nums))
    total := nums[0]
    for i := 1; i < len(nums); i++ {
       prefixs[i] = total
       total += nums[i]
    }
    return NumArray{
        nums,
        prefixs,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    // run at least 10^4 times => it need to be fast
    // Idea => cache some stuff so that the calculation can be fast
    return this.prefixs[right] + this.nums[right] - this.prefixs[left]
}
