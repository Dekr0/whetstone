package maxnumofballons

import "testing"

type TestCase struct {
   in string
   out int
}

func TestMaxNumOfBallons(t *testing.T) {
    test_cases := []TestCase{
        {"nlaebolko", 1},
        {"loonbalxballpoon", 2},
        {"leetcode", 0},
    } 
    for i, tc := range test_cases {
        r := MaxNumberOfBallons(tc.in)
        if r != tc.out {
            t.Fatalf("Test case %d failed with %v, expected %v", i, r, tc.out)
        }
    }
}
