package issubsequence

import "testing"

func TestIsSubequence(t *testing.T) {
    input := [][]string{
        { "abc", "ahbgdc" },
        { "axc", "ahbgdc" },
    }

    output := []bool{ true, false }

    for i := 0; i < len(input); i++ {
        if IsSubsequence(input[i][0], input[i][1]) != output[i] {
            t.Fatalf("Test Case %d failed", i)
        }
    }
}
