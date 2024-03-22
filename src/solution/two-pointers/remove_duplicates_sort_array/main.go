package main

func solutionA(nums []int) int {
    set := make(map[int]bool)
    j := -1;
    for i := 0; i < len(nums); i++ {
        _, in := set[nums[i]]
        if in {
            if (j == -1) {
                j = i
            }
        } else {
            set[nums[i]] = true
            if (j != -1) {
                nums[j] = nums[i]
                j++
            }
        }
    }
    return len(set)
}

func solutionB(nums []int) int {
    j := 0;
    /* Utilize the properties of the given input. This can help formulate a 
    solution or reduce the amount of state and condition check since these 
    properties provide some form of implication
    */
    for i := 0; i < len(nums); i++ {
        if nums[i] > nums[j] {
            nums[j + 1] = nums[i]
            j++
        }
    }
    return j + 1
}
