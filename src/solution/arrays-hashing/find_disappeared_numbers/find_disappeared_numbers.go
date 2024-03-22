package finddisappearednumbers


func abs(n int) int {
    if n >= 0 {
        return n
    }

    return -n
}


func FindDisappearedNumbers(nums []int) []int {
    /*
     * Hint for O(N) runtime and O(1) space
     *
     * How to achieve O(1) space? Use and modify input array to keep track of 
     * state.
     *
     * How to ensure a number appear in the array? General approach => flip 
     * the sign. Since for each input arary with length n, each element at 
     * index i is mapped to value i + 1 where 1 <= i <= n
     *
     * How can we filp the sign? It's rather untrivial. 
     * 1. We will take the abs. value of the value at index i,
     * 2. After taking the abs., use it as an index, flip the sign of the value 
     * at this index
     * 
     * You most likley ask step 2 will override the value (e.g. [4, 1, 1, 1])
     * This is why step 3 comes in. If one of the slot (say j) is greater zero, 
     * that's mean j + 1 is missing.
     * 
     * The key to understand this work is because if a number, k, appear in this 
     * array, that means the value at the index k - 1 is guarantee being < 0
    */ 
    for _, n := range nums {
        i := abs(n) - 1
        nums[i] = -1 * abs(nums[i])
    }

    res := []int{}

    for i, n := range nums {
        if n > 0 {
            res = append(res, i + 1)
        }
    }

    return res
}
