package twosum

func TwoSum(nums []int, target int) []int {
    remain := make(map[int]int)
    for i, n := range nums {
        v, in := remain[target - n]
        if in {
            return []int{i, v}
        }
        remain[n] = i
    }
    return []int{}
}
