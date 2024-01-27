package isomorphicstrings

import "testing"

func TestIsomorphicStrings(t *testing.T) {
    type TestCase struct {
        in []string
        out bool
    }

    test_cases := []TestCase {
        {
            []string{"egg", "add"},
            true,
        },
        {
            []string{"foo", "bar"},
            false,
        },
        {
            []string{"paper", "title"},
            true,
        },
        {
            []string{"badc", "baba"},
            false,
        },
    }

   for i, tc := range test_cases {
       r := IsomorphicStrings(tc.in[0], tc.in[1])
       if r != tc.out {
           t.Fatalf("Test case %d failed with %v", i, r)
       }
   }
}
