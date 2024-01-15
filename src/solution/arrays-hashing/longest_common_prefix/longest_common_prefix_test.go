package longestcommonprefix

import "testing"

type TestCase struct {
    in []string
    out string
}

func TestLongestCommonPrefix(t *testing.T) {
    test_cases := []TestCase {
        { []string{ "flower", "flow", "flight" }, "fl" },
        { []string{ "dog", "racecar", "car" }, "" },
    }
    for i, tc := range test_cases {
        if LongestCommonPrefix(tc.in) != tc.out {
            t.Fatalf("Test case %d failed", i)
        }
    }
}

func TestLongestCommonCryptic(t *testing.T) {
    test_cases := []TestCase {
        { []string{ "flower", "flow", "flight" }, "fl" },
        { []string{ "dog", "racecar", "car" }, "" },
    }
    for i, tc := range test_cases {
        if LongestCommonPrefixCryptic(tc.in) != tc.out {
            t.Fatalf("Test case %d failed", i)
        }
    }
}
