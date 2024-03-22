package main

func solutionA(nums1 []int, m int, nums2 []int, n int) {
    if (m == 0) {
        for ; n > 0; n-- {
            nums1[n - 1] = nums2[n - 1]
        }
        return
    }

    /* When thinking about state (reduce extra state) optimization, ask whether
     if a state can be derived from other states. In this case, total length can
     be from m + n. Since either m decrement or n decrement assuming they work 
     as the intended way, m + n - 1 is an valid index for nums1 []int
    */

    /* Ask why I have to use this approach when I get stuck on problem. Break 
    the current chain of thought and try a different way.
    */

    /* What's the easy way or less complex way to solve this problem as a human
    ? This can be one way of finding the solution, and the solution is more 
    intuitive. In this case, why shuffling everything in nums2 [] when one can 
    start at the end of nums1 [], find the max. for each index, and overwrite 
    nums1[]
    */
    for n > 0 {
        if m > 0 && nums1[m - 1] > nums2[n - 1] {
            nums1[m + n - 1] = nums1[m - 1]
            m--
        } else {
           nums1[m + n - 1] = nums2[n - 1] 
           n--
        }
    }
}

