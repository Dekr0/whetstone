package nextgreaterelement

import (
	"slices"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
    type TestCase struct {
        in_1 []int
        in_2 []int
        out []int
    }

    test_cases := []TestCase{
        {
            []int{4,1,2},
            []int{1,3,4,2},
            []int{-1,3,-1},
        },
        {
            []int{2,4},
            []int{1,2,3,4},
            []int{3,-1},
        },
    }

    for i, tc := range test_cases {
        if !slices.Equal(NextGreaterElement(tc.in_1, tc.in_2), tc.out) {
            t.Fatalf("Test %d failed", i)
        }
    }
}
