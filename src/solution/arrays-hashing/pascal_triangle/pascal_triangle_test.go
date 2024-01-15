package pascaltriangle

import "testing"

type TestCase struct {
    in int
    out [][]int
}

func TestGenerate(t *testing.T) {
    test_cases := []TestCase{
        { 5, [][]int{ {1}, {1,1}, {1,2,1}, {1,3,3,1}, {1,4,6,4,1}, } },
        { 1, [][]int{ {1}, } },
    }
    for _, tc := range test_cases {
        out := Generate(tc.in)
        t.Logf("%v\n", out)
    }
}
