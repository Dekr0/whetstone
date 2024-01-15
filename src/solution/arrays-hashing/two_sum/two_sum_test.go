package twosum

import (
	"slices"
	"testing"
)

type Test struct {
    nums []int
    target int
}

func TestTwoSum(t *testing.T) {
    in := []Test{
        {[]int{2,7,11,15},9},
        {[]int{3,2,4},6},
        {[]int{3,3},6},
    }
    out := [][]int{
        {0,1},
        {1,2},
        {0,1},
    }
    for i := 0; i < len(in); i++ {
        r := TwoSum(in[i].nums, in[i].target)
        if !slices.Contains(out[i], r[0]) || !slices.Contains(out[i], r[1]) {
            t.Fatalf("Test Case %d failed", i)
        }
    }
}
