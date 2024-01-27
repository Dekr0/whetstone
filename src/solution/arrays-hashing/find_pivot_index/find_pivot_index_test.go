package findpivotindex

import "testing"

func TestFindPivotIndex(t *testing.T) {
    type TestCase struct {
        in []int
        out int
    }

    test_cases := []TestCase{
        {
            []int{1,7,3,6,5,6},
            3,
        },
        {
            []int{1,2,3},
            -1,
        },
        {
            []int{2,1,-1},
            0,
        },
    }

    for _, tc := range test_cases {
        t.Log(FindPivotIndexOpt(tc.in))
    }
}
