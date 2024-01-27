package canplaceflowers

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
    type TestCase struct {
        bed []int
        n int
        out bool
    }

    test_cases := []TestCase {
        {
            []int{1, 0, 0, 0, 1}, 1, true,
        },
        {
            []int{1, 0, 0, 0, 1}, 2, false,
        },
        {
            []int{1}, 1, false,
        },
        {
            []int{1, 0}, 2, false,
        },
        {
            []int{0, 0, 0, 0, 0, 0, 0}, 4, true,
        },
        {
            []int{0, 0, 0, 0, 0}, 3, true,
        },
    }
    
   for i, tc := range test_cases {
       r := CanPlacePowerBetter(tc.bed, tc.n)
       if r != tc.out {
           t.Fatalf("Test case %d failed with %v", i, r)
       }
   }
}
