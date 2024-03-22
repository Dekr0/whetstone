package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
    dict := make(map[int]int)
    // O(n)
    for _, n := range nums {
        _, in := dict[n]
        if !in {
            dict[n] = 1
        } else {
            dict[n] += 1
        }
    }
    bucket := make([][]int, len(nums))
    i := 0
    for ; i < len(nums); i++ {
        bucket[i] = []int{}
    }
    for key, v := range dict {
        bucket[v] = append(bucket[v], key)
    }
    
    fmt.Printf("%v\n", bucket)
    return []int{}
}

func main() {
    topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
}
