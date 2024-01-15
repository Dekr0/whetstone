package uniqueemails

import "testing"

func TestNumUniqueEmails(t *testing.T) {
   type TestCase struct {
       in []string
       out int
   } 

   test_cases := []TestCase{
       {
           []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"},
           2,
       },
       {
           []string{"a@leetcode.com","b@leetcode.com","c@leetcode.com"},
           3,
       },
   }

   for i, tc := range test_cases {
       r := NumUniqueEmails(tc.in)
       if r != tc.out {
           t.Fatalf("Test case %d failed with %v", i, r)
       }
   }
}
