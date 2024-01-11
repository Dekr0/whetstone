package main

import (
    mapSet "github.com/deckarep/golang-set/v2"
)


func ContainsDuplicate(nums []int) bool {
    /*
    1 <= nums.length <= 10^5
    -10^9 <= nums[i] <= 10^9
    */

    set := make(map[int]Void)
    for _, n := range nums {
        _, exists := set[n]
        if exists {
            return true
        }
        set[n] = Void{}
    }
    return false
}

func ContainsDuplicateAlt1(nums []int) bool {
    set := mapSet.NewSet[int]()
    last_size := set.Cardinality()
    for _, n := range nums {
        set.Add(n)
        if last_size == set.Cardinality() {
            return true
        }
        last_size = set.Cardinality()
    }
    return false
}
