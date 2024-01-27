package majorityelements

import (
	"testing"
)

func TestMajorityElements(t *testing.T) {
    type TestCase struct {
        in []int
        out int
    }

    test_cases := []TestCase{
        {
            []int{3, 2, 3},
            3,
        },
        {
            []int{2, 2, 1, 1, 1, 2, 2},
            2,
        },
    }

    for i, tc := range test_cases {
        r := MajorityElements(tc.in) 
        if r != tc.out {
            t.Fatalf("Test Case %d failed with %d", i, r)
        }
    }
}
