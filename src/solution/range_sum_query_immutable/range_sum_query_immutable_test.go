package rangesumqueryimmutable

import "testing"


func Test(t *testing.T) {
    type TestCase struct {
        in []int
        out int
    }
    arr := []int{-2, 0, 3, -5, 2, -1}
    tests := []TestCase{
        {
            []int{0, 2},
            1,
        },
        {
            []int{2, 5},
            -1,
        },
        {
            []int{0, 5},
            -3,
        },
    }
    obj := Constructor(arr)
    t.Log(obj.prefixs)
    for i, tc := range tests {
        r := obj.SumRange(tc.in[0], tc.in[1])
        if r != tc.out {
            t.Fatalf("Test %d failed with %v", i, r)
        }
    }
}
